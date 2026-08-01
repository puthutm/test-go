"use client";

import flatpickr from "flatpickr";
import { useRef, useState } from "react";
import Flatpickr, { DateTimePickerProps } from "react-flatpickr";
import { Indonesian } from "flatpickr/dist/l10n/id.js";

import { classMerge } from "@/lib/utils/class-merge";

interface DatePickerProps extends DateTimePickerProps {
  onChange: (e: Date | Date[]) => void;
  options?: flatpickr.Options.Options | undefined;
  placeholder?: string;
  classNameClearIcon?: string;
  classNameFlatpickr?: string;
}

export const DatePicker: React.FC<DatePickerProps> = ({
  onChange,
  options,
  placeholder = "Pilih tanggal",
  className,
  classNameClearIcon,
  classNameFlatpickr,
  ...props
}) => {
  const [dates, setDates] = useState<Date[] | null>(null);

  const refDate = useRef<Flatpickr>(null);

  const clearDate = () => {
    refDate?.current?.flatpickr.clear(true);
    setDates(null);
  };

  return (
    <div
      className={classMerge(
        className,
        "position-relative form-control form-control-icon"
      )}
      style={{ zIndex: 0 }}
    >
      <Flatpickr
        className={classMerge(
          classNameFlatpickr,
          "w-100 h-100 rounded disabled-input"
        )}
        options={{ ...options, locale: Indonesian }}
        placeholder={placeholder}
        onChange={(e) => {
          onChange(e);
          setDates(e);
        }}
        ref={refDate}
        style={{ border: "none", zIndex: 40 }}
        {...props}
      />
      {dates?.length ? (
        <span
          className={`input-group-text bg-transparent border-0 position-absolute  ${
            classNameClearIcon ? classNameClearIcon : "bottom-25"
          }`}
          style={{
            cursor: "pointer",
            zIndex: 100,
            color: "#878A99",
            right: classNameClearIcon ? 0 : "16px",
          }}
          onClick={clearDate}
        >
          <i className="mdi mdi-close" />
        </span>
      ) : null}
    </div>
  );
};

{
  /**
  example use
  
  <Row className="align-items-center">
    <Col lg={4}>
      <Label
        htmlFor="number_of_study_program"
        className="form-label mb-1 required fw-semibold"
      >
        Tanggal Lahir
      </Label>
    </Col>
    <Col lg={8}>
      <div className="input-group w-100">
        <i
          className="pe-3 input-group-text bg-transparent"
          style={{ zIndex: 200 }}
        >
          <CalendarMonthIcon />
        </i>
        <DatePicker
          onChange={(e) => console.log(e)}
          options={{
            mode: "single",
            dateFormat: "d F Y",
          }}
        />
      </div>
    </Col>
  </Row>
  
  */
}
