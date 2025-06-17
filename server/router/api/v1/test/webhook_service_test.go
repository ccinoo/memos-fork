package v1

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	v1pb "github.com/usememos/memos/proto/gen/api/v1"
)

func TestCreateWebhook(t *testing.T) {
	ctx := context.Background()

	t.Run("CreateWebhook with host user", func(t *testing.T) {
		// Create test service for this specific test
		ts := NewTestService(t)
		defer ts.Cleanup()

		// Create and authenticate as host user
		hostUser, err := ts.CreateHostUser(ctx, "admin")
		require.NoError(t, err)

		userCtx := ts.CreateUserContext(ctx, hostUser.Username)

		// Create a webhook
		req := &v1pb.CreateWebhookRequest{
			Webhook: &v1pb.Webhook{
				DisplayName: "Test Webhook",
				Url:         "https://example.com/webhook",
			},
		}

		resp, err := ts.Service.CreateWebhook(userCtx, req)

		// Verify response
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, "Test Webhook", resp.DisplayName)
		require.Equal(t, "https://example.com/webhook", resp.Url)
		require.Contains(t, resp.Name, "webhooks/")
		require.Equal(t, fmt.Sprintf("users/%d", hostUser.ID), resp.Creator)
	})

	t.Run("CreateWebhook fails without authentication", func(t *testing.T) {
		// Create test service for this specific test
		ts := NewTestService(t)
		defer ts.Cleanup()

		// Try to create webhook without authentication
		req := &v1pb.CreateWebhookRequest{
			Webhook: &v1pb.Webhook{
				DisplayName: "Test Webhook",
				Url:         "https://example.com/webhook",
			},
		}

		_, err := ts.Service.CreateWebhook(ctx, req)

		// Should fail with permission denied or unauthenticated
		require.Error(t, err)
	})

	t.Run("CreateWebhook fails with regular user", func(t *testing.T) {
		// Create test service for this specific test
		ts := NewTestService(t)
		defer ts.Cleanup()

		// Create and authenticate as regular user
		regularUser, err := ts.CreateRegularUser(ctx, "user1")
		require.NoError(t, err)

		userCtx := ts.CreateUserContext(ctx, regularUser.Username)

		// Try to create webhook as regular user
		req := &v1pb.CreateWebhookRequest{
			Webhook: &v1pb.Webhook{
				DisplayName: "Test Webhook",
				Url:         "https://example.com/webhook",
			},
		}

		_, err = ts.Service.CreateWebhook(userCtx, req)

		// Should fail with permission denied
		require.Error(t, err)
		require.Contains(t, err.Error(), "permission denied")
	})

	t.Run("CreateWebhook validates required fields", func(t *testing.T) {
		// Create test service for this specific test
		ts := NewTestService(t)
		defer ts.Cleanup()

		// Create and authenticate as host user
		hostUser, err := ts.CreateHostUser(ctx, "admin")
		require.NoError(t, err)

		userCtx := ts.CreateUserContext(ctx, hostUser.Username)

		// Try to create webhook with missing URL
		req := &v1pb.CreateWebhookRequest{
			Webhook: &v1pb.Webhook{
				DisplayName: "Test Webhook",
				// URL missing
			},
		}

		_, err = ts.Service.CreateWebhook(userCtx, req)

		// Should fail with validation error
		require.Error(t, err)
	})
}

func TestListWebhooks(t *testing.T) {
	ctx := context.Background()

	t.Run("ListWebhooks returns empty list initially", func(t *testing.T) {
		// Create test service for this specific test
		ts := NewTestService(t)
		defer ts.Cleanup()

		// Create host user for authentication
		hostUser, err := ts.CreateHostUser(ctx, "admin")
		require.NoError(t, err)

		userCtx := ts.CreateUserContext(ctx, hostUser.Username)

		// List webhooks
		req := &v1pb.ListWebhooksRequest{}
		resp, err := ts.Service.ListWebhooks(userCtx, req)

		// Verify response
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Empty(t, resp.Webhooks)
	})

	t.Run("ListWebhooks returns created webhooks", func(t *testing.T) {
		// Create test service for this specific test
		ts := NewTestService(t)
		defer ts.Cleanup()

		// Create host user and authenticate
		hostUser, err := ts.CreateHostUser(ctx, "admin")
		require.NoError(t, err)
		userCtx := ts.CreateUserContext(ctx, hostUser.Username)

		// Create a webhook
		createReq := &v1pb.CreateWebhookRequest{
			Webhook: &v1pb.Webhook{
				DisplayName: "Test Webhook",
				Url:         "https://example.com/webhook",
			},
		}
		createdWebhook, err := ts.Service.CreateWebhook(userCtx, createReq)
		require.NoError(t, err)

		// List webhooks
		listReq := &v1pb.ListWebhooksRequest{}
		resp, err := ts.Service.ListWebhooks(userCtx, listReq)

		// Verify response
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Len(t, resp.Webhooks, 1)
		require.Equal(t, createdWebhook.Name, resp.Webhooks[0].Name)
		require.Equal(t, createdWebhook.Url, resp.Webhooks[0].Url)
	})

	t.Run("ListWebhooks fails without authentication", func(t *testing.T) {
		// Create test service for this specific test
		ts := NewTestService(t)
		defer ts.Cleanup()

		// Try to list webhooks without authentication
		req := &v1pb.ListWebhooksRequest{}
		_, err := ts.Service.ListWebhooks(ctx, req)

		// Should fail with permission denied or unauthenticated
		require.Error(t, err)
	})
}

func TestGetWebhook(t *testing.T) {
	ctx := context.Background()

	t.Run("GetWebhook returns webhook by name", func(t *testing.T) {
		// Create test service for this specific test
		ts := NewTestService(t)
		defer ts.Cleanup()

		// Create host user and authenticate
		hostUser, err := ts.CreateHostUser(ctx, "admin")
		require.NoError(t, err)
		userCtx := ts.CreateUserContext(ctx, hostUser.Username)

		// Create a webhook
		createReq := &v1pb.CreateWebhookRequest{
			Webhook: &v1pb.Webhook{
				DisplayName: "Test Webhook",
				Url:         "https://example.com/webhook",
			},
		}
		createdWebhook, err := ts.Service.CreateWebhook(userCtx, createReq)
		require.NoError(t, err)

		// Get the webhook
		getReq := &v1pb.GetWebhookRequest{
			Name: createdWebhook.Name,
		}
		resp, err := ts.Service.GetWebhook(userCtx, getReq)

		// Verify response
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, createdWebhook.Name, resp.Name)
		require.Equal(t, createdWebhook.Url, resp.Url)
		require.Equal(t, createdWebhook.Creator, resp.Creator)
	})

	t.Run("GetWebhook fails with invalid name", func(t *testing.T) {
		// Create test service for this specific test
		ts := NewTestService(t)
		defer ts.Cleanup()

		// Create host user and authenticate
		hostUser, err := ts.CreateHostUser(ctx, "admin")
		require.NoError(t, err)
		userCtx := ts.CreateUserContext(ctx, hostUser.Username)

		// Try to get webhook with invalid name
		req := &v1pb.GetWebhookRequest{
			Name: "invalid/webhook/name",
		}
		_, err = ts.Service.GetWebhook(userCtx, req)

		// Should return an error
		require.Error(t, err)
	})

	t.Run("GetWebhook fails with non-existent webhook", func(t *testing.T) {
		// Create test service for this specific test
		ts := NewTestService(t)
		defer ts.Cleanup()

		// Create host user and authenticate
		hostUser, err := ts.CreateHostUser(ctx, "admin")
		require.NoError(t, err)
		userCtx := ts.CreateUserContext(ctx, hostUser.Username)

		// Try to get non-existent webhook
		req := &v1pb.GetWebhookRequest{
			Name: "webhooks/999",
		}
		_, err = ts.Service.GetWebhook(userCtx, req)

		// Should return not found error
		require.Error(t, err)
		require.Contains(t, err.Error(), "not found")
	})
}

func TestUpdateWebhook(t *testing.T) {
	ctx := context.Background()

	t.Run("UpdateWebhook updates webhook properties", func(t *testing.T) {
		// Create test service for this specific test
		ts := NewTestService(t)
		defer ts.Cleanup()

		// Create host user and authenticate
		hostUser, err := ts.CreateHostUser(ctx, "admin")
		require.NoError(t, err)
		userCtx := ts.CreateUserContext(ctx, hostUser.Username)

		// Create a webhook
		createReq := &v1pb.CreateWebhookRequest{
			Webhook: &v1pb.Webhook{
				Name: "Original Webhook",
				Url:  "https://example.com/webhook",
			},
		}
		createdWebhook, err := ts.Service.CreateWebhook(userCtx, createReq)
		require.NoError(t, err)

		// Update the webhook
		updateReq := &v1pb.UpdateWebhookRequest{
			Webhook: &v1pb.Webhook{
				Name: createdWebhook.Name,
				Url:  "https://updated.example.com/webhook",
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{"url"},
			},
		}
		resp, err := ts.Service.UpdateWebhook(userCtx, updateReq)

		// Verify response
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, createdWebhook.Name, resp.Name)
		require.Equal(t, "https://updated.example.com/webhook", resp.Url)
	})

	t.Run("UpdateWebhook fails without authentication", func(t *testing.T) {
		// Create test service for this specific test
		ts := NewTestService(t)
		defer ts.Cleanup()

		// Try to update webhook without authentication
		req := &v1pb.UpdateWebhookRequest{
			Webhook: &v1pb.Webhook{
				Name: "webhooks/1",
				Url:  "https://updated.example.com/webhook",
			},
		}

		_, err := ts.Service.UpdateWebhook(ctx, req)

		// Should fail with permission denied or unauthenticated
		require.Error(t, err)
	})
}

func TestDeleteWebhook(t *testing.T) {
	ctx := context.Background()

	t.Run("DeleteWebhook removes webhook", func(t *testing.T) {
		// Create test service for this specific test
		ts := NewTestService(t)
		defer ts.Cleanup()

		// Create host user and authenticate
		hostUser, err := ts.CreateHostUser(ctx, "admin")
		require.NoError(t, err)
		userCtx := ts.CreateUserContext(ctx, hostUser.Username)

		// Create a webhook
		createReq := &v1pb.CreateWebhookRequest{
			Webhook: &v1pb.Webhook{
				DisplayName: "Test Webhook",
				Url:         "https://example.com/webhook",
			},
		}
		createdWebhook, err := ts.Service.CreateWebhook(userCtx, createReq)
		require.NoError(t, err)

		// Delete the webhook
		deleteReq := &v1pb.DeleteWebhookRequest{
			Name: createdWebhook.Name,
		}
		_, err = ts.Service.DeleteWebhook(userCtx, deleteReq)

		// Verify deletion
		require.NoError(t, err)

		// Try to get the deleted webhook
		getReq := &v1pb.GetWebhookRequest{
			Name: createdWebhook.Name,
		}
		_, err = ts.Service.GetWebhook(userCtx, getReq)

		// Should return not found error
		require.Error(t, err)
		require.Contains(t, err.Error(), "not found")
	})

	t.Run("DeleteWebhook fails without authentication", func(t *testing.T) {
		// Create test service for this specific test
		ts := NewTestService(t)
		defer ts.Cleanup()

		// Try to delete webhook without authentication
		req := &v1pb.DeleteWebhookRequest{
			Name: "webhooks/1",
		}

		_, err := ts.Service.DeleteWebhook(ctx, req)

		// Should fail with permission denied or unauthenticated
		require.Error(t, err)
	})

	t.Run("DeleteWebhook fails with non-existent webhook", func(t *testing.T) {
		// Create test service for this specific test
		ts := NewTestService(t)
		defer ts.Cleanup()

		// Create host user and authenticate
		hostUser, err := ts.CreateHostUser(ctx, "admin")
		require.NoError(t, err)
		userCtx := ts.CreateUserContext(ctx, hostUser.Username)

		// Try to delete non-existent webhook
		req := &v1pb.DeleteWebhookRequest{
			Name: "webhooks/999",
		}
		_, err = ts.Service.DeleteWebhook(userCtx, req)

		// Should return not found error
		require.Error(t, err)
		require.Contains(t, err.Error(), "not found")
	})
}
