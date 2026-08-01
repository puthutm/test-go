"use client";

import { AutoFocusPlugin } from "@lexical/react/LexicalAutoFocusPlugin";
import { ContentEditable } from "@lexical/react/LexicalContentEditable";
import { LexicalErrorBoundary } from "@lexical/react/LexicalErrorBoundary";
import { HistoryPlugin } from "@lexical/react/LexicalHistoryPlugin";
import { RichTextPlugin } from "@lexical/react/LexicalRichTextPlugin";
import { OnChangePlugin } from "@lexical/react/LexicalOnChangePlugin";
import { TabIndentationPlugin } from "@lexical/react/LexicalTabIndentationPlugin";
import { ClickableLinkPlugin } from "@lexical/react/LexicalClickableLinkPlugin";
import { useLexicalEditable } from "@lexical/react/useLexicalEditable";
import { useCallback, useState } from "react";
import { EditorState, LexicalEditor } from "lexical";
import { ListPlugin } from "@lexical/react/LexicalListPlugin";
import { CheckListPlugin } from "@lexical/react/LexicalCheckListPlugin";

import ToolbarPlugin from "./plugins/toolbar";

import "./style.css";
import LinkPlugin from "./plugins/link-plugin";
import FloatingLinkEditorPlugin from "./plugins/floating-link-editor";
import { ToggleEditablePlugin } from "./plugins/toogle-editable-plugin";
import { SetEditorValuePlugin } from "./plugins/set-editor-value-plugin";
import LexicalAutoLinkPlugin from "./plugins/auto-link-plugin";
import CodeHighlightPlugin from "./plugins/code-highlight-plugin";

interface EditorProps {
  onChange: (
    editorState: EditorState,
    editor?: LexicalEditor,
    tags?: Set<string>
  ) => void;
  value?: string;
  disabled?: boolean;
  placeholder?: string;
  isError?: boolean;
}

export default function Editor({
  value,
  onChange,
  disabled,
  placeholder = "Masukkan teks di sini",
  isError,
}: EditorProps) {
  const [floatingAnchorElem, setFloatingAnchorElem] =
    useState<HTMLDivElement | null>(null);
  const [isLinkEditMode, setIsLinkEditMode] = useState<boolean>(false);
  const isEditable = useLexicalEditable();

  const onRef = useCallback(
    (node: HTMLDivElement | null) => {
      if (node !== null && node !== floatingAnchorElem) {
        setFloatingAnchorElem(node);
      }
    },
    [floatingAnchorElem]
  );

  return (
    <div className="editor-container">
      <ToolbarPlugin setIsLinkEditMode={setIsLinkEditMode} />
      <div className={`editor-inner ${isError ? "border border-danger" : ""}`}>
        <RichTextPlugin
          contentEditable={
            <div className="editor-scroller">
              <div className="editor" ref={onRef}>
                <ContentEditable className="editor-input" />
              </div>
            </div>
          }
          placeholder={<div className="editor-placeholder">{placeholder}</div>}
          ErrorBoundary={LexicalErrorBoundary}
        />
        <HistoryPlugin />
        <AutoFocusPlugin />
        {floatingAnchorElem && (
          <FloatingLinkEditorPlugin
            anchorElem={floatingAnchorElem}
            isLinkEditMode={isLinkEditMode}
            setIsLinkEditMode={setIsLinkEditMode}
          />
        )}
        <LexicalAutoLinkPlugin />
        <OnChangePlugin
          onChange={(state) => {
            onChange(state);
          }}
        />
        <LinkPlugin hasLinkAttributes={true} />
        <CodeHighlightPlugin />
        <ClickableLinkPlugin disabled={isEditable} />
        <SetEditorValuePlugin value={value} />
        <TabIndentationPlugin maxIndent={7} />
        <ToggleEditablePlugin disabled={disabled} />
        <ListPlugin hasStrictIndent={false} />
        <CheckListPlugin />
      </div>
    </div>
  );
}
