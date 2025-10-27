# Memos

<img align="right" height="96px" src="https://raw.githubusercontent.com/usememos/.github/refs/heads/main/assets/logo-rounded.png" alt="Memos" />

An open-source, self-hosted note-taking services. Write, organize, and own your notes. No tracking, no ads, no subscription fees — just a clean, fast way to capture your ideas.

[![Home](https://img.shields.io/badge/🏠-usememos.com-blue?style=flat-square)](https://www.usememos.com)
[![Live Demo](https://img.shields.io/badge/✨-Try%20Demo-orange?style=flat-square)](https://demo.usememos.com/)
[![Docs](https://img.shields.io/badge/📚-Documentation-green?style=flat-square)](https://www.usememos.com/docs)
[![Discord](https://img.shields.io/badge/💬-Discord-5865f2?style=flat-square&logo=discord&logoColor=white)](https://discord.gg/tfPJa4UmAv)
[![Docker Pulls](https://img.shields.io/docker/pulls/neosmemo/memos?style=flat-square&logo=docker)](https://hub.docker.com/r/neosmemo/memos)

<img src="https://raw.githubusercontent.com/usememos/.github/refs/heads/main/assets/demo.png" alt="Memos Demo Screenshot" height="512" />

### 💎 Featured Sponsors

[**Warp** — The AI-powered terminal built for speed and collaboration](https://go.warp.dev/memos)

<a href="https://go.warp.dev/memos" target="_blank" rel="noopener">
  <img src="https://raw.githubusercontent.com/warpdotdev/brand-assets/main/Github/Sponsor/Warp-Github-LG-02.png" alt="Warp - The terminal for the 21st century" height="300" />
</a>

## 🎯 Why Memos?

**Your thoughts. Your data. Your control.**

Memos is an open-source, self-hosted alternative to cloud note-taking services. No tracking, no ads, no subscription fees — just a clean, fast way to capture and organize your ideas.

### 🆚 What Makes Memos Different?

| Feature           | Memos                          | Cloud Services                |
| ----------------- | ------------------------------ | ----------------------------- |
| **Privacy**       | ✅ Self-hosted, zero telemetry | ❌ Your data on their servers |
| **Cost**          | ✅ Free forever, MIT license   | ❌ Subscription fees          |
| **Performance**   | ✅ Instant load, no latency    | ⚠️ Depends on internet        |
| **Ownership**     | ✅ Full control & export       | ❌ Vendor lock-in             |
| **API Access**    | ✅ Full REST + gRPC APIs       | ⚠️ Limited or paid            |
| **Customization** | ✅ Open source, forkable       | ❌ Closed ecosystem           |
| **Deployment**    | ✅ Single binary, Docker       | ⚠️ SaaS only                  |

## ✨ Features

- **🔒 Privacy-First Architecture** — Self-hosted on your infrastructure with no telemetry or tracking. Your data stays yours with full ownership and export capabilities.

- **⚡ Lightning-Fast Performance** — Built with Go and React for exceptional speed. Real-time updates ensure your notes are always in sync across devices.

- **📝 Markdown Native** — Write naturally with full markdown support, code highlighting, and rich text rendering. No lock-in, just plain text you can take anywhere.

- **🐳 Deploy in Seconds** — One Docker command to get started. Simple deployment with Docker Compose, binaries, or Kubernetes.

- **🎨 Beautiful & Minimal UI** — Clean, intuitive interface that stays out of your way. Mobile-responsive design with dark mode support.

- **🔗 Powerful API** — RESTful and gRPC APIs for seamless integration with your existing workflow. Build custom tools and automations easily.

- **🌍 Open Source & Community-Driven** — MIT licensed with an active community. Regular updates and transparent development on GitHub.

## 💡 Perfect For

- 📓 **Personal Journaling** — Capture daily thoughts, reflections, and life moments in a private space
- 🧠 **Knowledge Management** — Build your second brain with interconnected notes and tags
- 📋 **Quick Captures** — Instantly jot down ideas before they vanish
- 💻 **Developer Notes** — Store code snippets, commands, and technical references with syntax highlighting
- 🔗 **Link Collections** — Bookmark and organize articles, resources, and inspiration
- 👥 **Team Wikis** — Share knowledge across your team with collaborative docs
- 🎓 **Learning Logs** — Track your learning journey, courses, and new skills
- 📝 **Meeting Notes** — Keep all your meeting agendas and action items organized

## 🚀 Quick Start

Get up and running in **under 30 seconds**:

```bash
docker run -d \
  --name memos \
  -p 5230:5230 \
  -v ~/.memos:/var/opt/memos \
  neosmemo/memos:stable
```

Open `http://localhost:5230` and start writing! 🎉

### 🎮 Try the Live Demo

Don't want to install yet? Try our [live demo](https://demo.usememos.com/) first!

**Demo Credentials:**

- Username: `demo`
- Password: `secret`

### 📦 More Deployment Options

**Need more options?** Check out our [installation guide](https://www.usememos.com/docs/installation) for:

- Docker Compose (recommended for production)
- Pre-built binaries (Linux, macOS, Windows)
- Kubernetes deployments
- Building from source

## 💖 Support Memos

Love Memos? Help us keep it growing!

<a href="https://github.com/sponsors/usememos" target="_blank">
  <img src="https://img.shields.io/badge/❤️_Sponsor_on_GitHub-ea4aaa?style=for-the-badge&logo=github-sponsors&logoColor=white" alt="Sponsor on GitHub">
</a>

**Community Sponsors:**

<a href="https://github.com/yourselfhosted" target="_blank"><img src="https://avatars.githubusercontent.com/u/140182318?v=4" alt="yourselfhosted" height="50" style="border-radius: 50%; margin: 5px;" /></a>
<a href="https://github.com/fixermark" target="_blank"><img src="https://avatars.githubusercontent.com/u/169982?v=4" alt="fixermark" height="50" style="border-radius: 50%; margin: 5px;" /></a>
<a href="https://github.com/alik-agaev" target="_blank"><img src="https://avatars.githubusercontent.com/u/2662697?v=4" alt="alik-agaev" height="50" style="border-radius: 50%; margin: 5px;" /></a>

_Every contribution, big or small, makes a difference!_

## 🤝 Contributing

We love contributions! Whether you're fixing bugs, adding features, improving docs, or spreading the word — every contribution matters.

**Get involved:**

- 🐛 [Report bugs](https://github.com/usememos/memos/issues/new?template=bug_report.md)
- 💡 [Suggest features](https://github.com/usememos/memos/issues/new?template=feature_request.md)
- 🔧 [Submit pull requests](https://github.com/usememos/memos/pulls)
- 📖 [Improve documentation](https://github.com/usememos/memos/tree/main/docs)
- 🌍 [Help with translations](https://github.com/usememos/memos/tree/main/web/src/locales)

## ⭐ Star History

[![Star History Chart](https://api.star-history.com/svg?repos=usememos/memos&type=Date)](https://star-history.com/#usememos/memos&Date)

---

<div align="center">

**[Website](https://www.usememos.com)** •
**[Docs](https://www.usememos.com/docs)** •
**[Demo](https://demo.usememos.com/)** •
**[X](https://x.com/usememos)** •
**[Discord](https://discord.gg/tfPJa4UmAv)**

Made with ❤️ by the Memos community

**If you like Memos, give us a ⭐ on GitHub!**

</div>
