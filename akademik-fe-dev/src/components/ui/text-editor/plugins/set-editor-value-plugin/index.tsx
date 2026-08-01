"use client";

import { useLexicalComposerContext } from "@lexical/react/LexicalComposerContext";
import { SerializedEditorState } from "lexical";
import { useEffect } from "react";

export function SetEditorValuePlugin({
  value,
}: {
  value?: string | SerializedEditorState;
}) {
  const [editor] = useLexicalComposerContext();

  useEffect(() => {
    if (!value) return;

    let validValue: SerializedEditorState | null = null;

    try {
      validValue = typeof value === "string" ? JSON.parse(value) : value;
    } catch (err) {
      console.warn("Invalid JSON string passed as editor state:", err);
      return;
    }

    if (
      !validValue ||
      typeof validValue !== "object" ||
      !("root" in validValue) ||
      typeof (validValue as any).root !== "object" ||
      (validValue as any).root.type !== "root"
    ) {
      return;
    }

    editor.update(() => {
      try {
        const currentEditorState = editor.getEditorState();
        const newEditorState = editor.parseEditorState(
          validValue as SerializedEditorState
        );

        const currentJSON = JSON.stringify(currentEditorState.toJSON());
        const newJSON = JSON.stringify(newEditorState.toJSON());

        if (currentJSON !== newJSON) {
          editor.setEditorState(newEditorState);
        }
      } catch (err) {
        console.warn("Failed to parse Lexical editor state:", err);
      }
    });
  }, [editor, value]);

  return null;
}
