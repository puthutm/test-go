"use client";

import { blockTypeToBlockName } from "@/lib/contexts/toolbar-context";
import { LexicalEditor } from "lexical";

import DropDown, {
  dropDownActiveClass,
  DropDownItem,
} from "@/components/ui/text-editor/components/dropdown";
import {
  formatBulletList,
  formatCode,
  formatHeading,
  formatNumberedList,
  formatParagraph,
  formatQuote,
} from "@/components/ui/text-editor/plugins/utils";

export function BlockFormatDropDown({
  editor,
  blockType,
  disabled = false,
}: {
  blockType: keyof typeof blockTypeToBlockName;
  editor: LexicalEditor;
  disabled?: boolean;
}): JSX.Element {
  return (
    <DropDown
      disabled={disabled}
      buttonClassName="toolbar-item block-controls"
      buttonIconClassName={"icon block-type " + blockType}
      buttonLabel={blockTypeToBlockName[blockType]}
      buttonAriaLabel="Formatting options for text style"
    >
      <DropDownItem
        className={
          "item wide " + dropDownActiveClass(blockType === "paragraph")
        }
        onClick={() => formatParagraph(editor)}
      >
        <div className="icon-text-container">
          <i className="icon paragraph" />
          <span className="text">Normal</span>
        </div>
      </DropDownItem>
      <DropDownItem
        className={"item wide " + dropDownActiveClass(blockType === "h1")}
        onClick={() => formatHeading(editor, blockType, "h1")}
      >
        <div className="icon-text-container">
          <i className="icon h1" />
          <span className="text">Heading 1</span>
        </div>
      </DropDownItem>
      <DropDownItem
        className={"item wide " + dropDownActiveClass(blockType === "h2")}
        onClick={() => formatHeading(editor, blockType, "h2")}
      >
        <div className="icon-text-container">
          <i className="icon h2" />
          <span className="text">Heading 2</span>
        </div>
      </DropDownItem>
      <DropDownItem
        className={"item wide " + dropDownActiveClass(blockType === "h3")}
        onClick={() => formatHeading(editor, blockType, "h3")}
      >
        <div className="icon-text-container">
          <i className="icon h3" />
          <span className="text">Heading 3</span>
        </div>
      </DropDownItem>
      <DropDownItem
        className={"item wide " + dropDownActiveClass(blockType === "number")}
        onClick={() => formatNumberedList(editor, blockType)}
      >
        <div className="icon-text-container">
          <i className="icon numbered-list" />
          <span className="text">Numbered List</span>
        </div>
      </DropDownItem>
      <DropDownItem
        className={"item wide " + dropDownActiveClass(blockType === "bullet")}
        onClick={() => formatBulletList(editor, blockType)}
      >
        <div className="icon-text-container">
          <i className="icon bullet-list" />
          <span className="text">Bullet List</span>
        </div>
      </DropDownItem>
      <DropDownItem
        className={"item wide " + dropDownActiveClass(blockType === "quote")}
        onClick={() => formatQuote(editor, blockType)}
      >
        <div className="icon-text-container">
          <i className="icon quote" />
          <span className="text">Quote</span>
        </div>
      </DropDownItem>
      <DropDownItem
        className={"item wide " + dropDownActiveClass(blockType === "code")}
        onClick={() => formatCode(editor, blockType)}
      >
        <div className="icon-text-container">
          <i className="icon code" />
          <span className="text">Code Block</span>
        </div>
      </DropDownItem>
    </DropDown>
  );
}
