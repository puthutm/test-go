"use client";

import { forwardRef } from "react";
import { FieldError } from "react-hook-form";
import Select, { InputActionMeta } from "react-select";
import makeAnimated from "react-select/animated";
import CreatableSelect from "react-select/creatable";

import { classMerge } from "@/lib/utils/class-merge";
import styles from "@/styles/select.module.css";

const controlStyles = {
  base: styles.control_base,
  focus: styles.control_focus,
  nonFocus: styles.control_base,
  isDisabled: styles.control_disabled,
};
const containerStyles = styles.container;
const placeholderStyles = styles.placeholder;
const selectInputStyles = "text-foreground text-sm ml-3";
const valueContainerStyles = "text-foreground text-sm flex gap-1 py-1";
const singleValueStyles = "ml-3";
const multiValueStyles = styles.multi_value;
const multiValueLabelStyles = styles.multi_value_label;
const multiValueRemoveStyles = styles.multi_value_remove;
const indicatorsContainerStyles =
  "pr-1 gap-1 bg-background rounded-r-lg rounded";
const clearIndicatorStyles = styles.clear_indicator;
const indicatorSeparatorStyles = "";
const dropdownIndicatorStyles = "";
const menuStyles = styles.menu;
const optionsStyle = styles.options;
const groupHeadingStyles = "ml-3 mt-2 mb-1 text-gray-500 text-sm bg-background";
const noOptionsMessageStyles = "text-muted-foreground bg-background py-2";

interface SelectComponentProps {
  options: { label: string; value: string }[];
  value?: any;
  onChange?: (value: any) => void;
  onInputChange?: (newValue: string, actionMeta?: InputActionMeta) => void;
  isMulti?: boolean;
  isDisabled?: boolean;
  isClearable?: boolean;
  isLoading?: boolean;
  createAble?: boolean;
  filterOption?:
    | ((option: any, inputValue: string) => boolean)
    | null
    | undefined;
  placeholder?: string;
  isError?: boolean | FieldError | any;
  hasIcon?: boolean;
  id: any;
  onMenuScrollToBottom?: ((event: WheelEvent | TouchEvent) => void) | undefined;
  menuIsOpen?: boolean;
  hideIndicator?: boolean;
}

export const SelectComponent = forwardRef<any, SelectComponentProps>(
  (
    {
      options,
      value,
      onChange,
      isMulti,
      isDisabled,
      isLoading,
      createAble,
      placeholder,
      isError,
      hasIcon,
      isClearable,
      onInputChange,
      filterOption,
      id,
      onMenuScrollToBottom,
      menuIsOpen,
      hideIndicator,
    },
    ref
  ) => {
    const animatedComponents = makeAnimated();
    const Comp = createAble ? CreatableSelect : Select;
    return (
      <Comp
        ref={ref}
        unstyled
        isSearchable
        value={value}
        isDisabled={isDisabled}
        isMulti={isMulti}
        isClearable={isClearable}
        isLoading={isLoading}
        placeholder={placeholder}
        components={animatedComponents}
        defaultValue={value}
        options={options}
        noOptionsMessage={() => "Tidak ada opsi"}
        onChange={onChange}
        onInputChange={onInputChange}
        filterOption={filterOption}
        classNames={{
          control: ({ isFocused, isDisabled }) =>
            classMerge(
              isFocused ? controlStyles.focus : controlStyles.nonFocus,
              controlStyles.base,
              isError && `${controlStyles.base} border-danger`,
              hasIcon && `${controlStyles.base} ${styles.icon_padding}`,
              isDisabled && controlStyles.isDisabled
            ),
          placeholder: () => placeholderStyles,
          input: () => selectInputStyles,
          option: () => optionsStyle,
          menu: () => menuStyles,
          valueContainer: () => valueContainerStyles,
          singleValue: () => singleValueStyles,
          multiValue: () => multiValueStyles,
          multiValueLabel: () => multiValueLabelStyles,
          multiValueRemove: () => multiValueRemoveStyles,
          indicatorsContainer: () =>
            `${hideIndicator ? "d-none" : ""} ${indicatorsContainerStyles}`,
          clearIndicator: () => clearIndicatorStyles,
          indicatorSeparator: () => indicatorSeparatorStyles,
          dropdownIndicator: () => dropdownIndicatorStyles,
          groupHeading: () => groupHeadingStyles,
          noOptionsMessage: () => noOptionsMessageStyles,
          container: () => containerStyles,
          loadingMessage: () => "py-2",
        }}
        maxMenuHeight={180}
        id={id}
        instanceId={id}
        onMenuScrollToBottom={onMenuScrollToBottom}
        menuIsOpen={menuIsOpen}
      />
    );
  }
);

SelectComponent.displayName = "SelectComponent";
