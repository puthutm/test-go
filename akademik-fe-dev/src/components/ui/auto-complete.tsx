"use client";

import { ChangeEvent, useEffect, useRef, useState } from "react";
import { useDebouncedCallback } from "use-debounce";
import { Spinner } from "reactstrap";
import { FormErrorMessage } from "./form-error-message";

export interface OptionType {
  label: string;
  value: string | number;
  [key: string]: any;
}

interface AutoCompleteInputProps<T> {
  id: keyof T;
  data: OptionType[] | undefined;
  placeholder?: string;
  isLoading?: boolean;
  disabled?: boolean;
  errors?: any;
  register: any;
  setValue: any;
  clearErrors: any;
  onSearch?: (query: string) => void;
}

export const AutoCompleteInput = <T extends Record<string, any>>({
  id,
  data = [],
  placeholder,
  isLoading = false,
  disabled = false,
  errors,
  register,
  setValue,
  clearErrors,
  onSearch,
}: AutoCompleteInputProps<T>) => {
  const [input, setInput] = useState("");
  const [showDropdown, setShowDropdown] = useState(false);
  const inputRef = useRef<HTMLInputElement>(null);
  const wrapperRef = useRef<HTMLDivElement>(null);

  const debouncedSearch = useDebouncedCallback((val: string) => {
    onSearch?.(val);
  }, 500);

  useEffect(() => {
    if (input) {
      setShowDropdown(true);
    } else {
      setShowDropdown(false);
    }
  }, [input]);

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (
        wrapperRef.current &&
        !wrapperRef.current.contains(event.target as Node)
      ) {
        setShowDropdown(false);
      }
    };

    document.addEventListener("mousedown", handleClickOutside);
    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, []);

  const handleSelect = (option: OptionType) => {
    setValue(id, option.value as T[keyof T], { shouldValidate: true });
    clearErrors(id as string);
    setInput(option.label);
    setShowDropdown(false);
  };

  return (
    <div className="form-icon position-relative" ref={wrapperRef}>
      <input type="hidden" {...register(id)} />
      <input
        type="text"
        className={`form-control ${errors ? "border border-danger" : ""}`}
        placeholder={placeholder}
        value={input}
        onChange={(e: ChangeEvent<HTMLInputElement>) => {
          const val = e.target.value;
          setInput(val);
          debouncedSearch(val);
        }}
        ref={inputRef}
        autoComplete="off"
        disabled={disabled}
        onFocus={() => setShowDropdown(true)}
      />

      <FormErrorMessage errors={errors} />

      {isLoading && (
        <Spinner
          size="sm"
          style={{ position: "absolute", top: 10, right: 8 }}
        />
      )}

      {showDropdown && data?.length ? (
        <div
          style={{ maxHeight: 200, overflowY: "auto", zIndex: 99 }}
          className="position-absolute bg-white shadow border w-100 mt-1 rounded"
        >
          {data.map((item) => (
            <span
              key={item.value}
              className="w-100 d-block p-2 list-auto-complete"
              style={{ cursor: "pointer" }}
              onMouseDown={() => handleSelect(item)}
            >
              {item.label}
            </span>
          ))}
        </div>
      ) : null}
    </div>
  );
};
