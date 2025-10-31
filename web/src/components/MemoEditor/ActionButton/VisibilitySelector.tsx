import { CheckIcon, ChevronDownIcon } from "lucide-react";
import VisibilityIcon from "@/components/VisibilityIcon";
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger } from "@/components/ui/dropdown-menu";
import { Visibility } from "@/types/proto/api/v1/memo_service";
import { useTranslate } from "@/utils/i18n";

interface Props {
  value: Visibility;
  onChange: (visibility: Visibility) => void;
  onOpenChange?: (open: boolean) => void;
}

const VisibilitySelector = (props: Props) => {
  const { value, onChange } = props;
  const t = useTranslate();

  const visibilityOptions = [
    { value: Visibility.PRIVATE, label: t("memo.visibility.private") },
    { value: Visibility.PROTECTED, label: t("memo.visibility.protected") },
    { value: Visibility.PUBLIC, label: t("memo.visibility.public") },
  ];

  const currentLabel = visibilityOptions.find((option) => option.value === value)?.label || "";

  return (
    <DropdownMenu onOpenChange={props.onOpenChange}>
      <DropdownMenuTrigger asChild>
        <button className="inline-flex items-center gap-1.5 px-2 text-sm text-muted-foreground opacity-80 hover:opacity-100 transition-colors">
          <VisibilityIcon visibility={value} className="opacity-60" />
          <span>{currentLabel}</span>
          <ChevronDownIcon className="w-3 h-3 opacity-60" />
        </button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end">
        {visibilityOptions.map((option) => (
          <DropdownMenuItem key={option.value} className="cursor-pointer gap-2" onClick={() => onChange(option.value)}>
            <VisibilityIcon visibility={option.value} />
            <span className="flex-1">{option.label}</span>
            {value === option.value && <CheckIcon className="w-4 h-4 text-primary" />}
          </DropdownMenuItem>
        ))}
      </DropdownMenuContent>
    </DropdownMenu>
  );
};

export default VisibilitySelector;
