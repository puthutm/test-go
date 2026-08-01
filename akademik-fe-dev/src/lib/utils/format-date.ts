/** result : "1990-01-12" */
export const formatDateNumeric = (dateString: string | number): string => {
  const date = new Date(dateString);
  return new Intl.DateTimeFormat("id-ID", {
    year: "numeric",
    month: "2-digit",
    day: "2-digit",
  })
    .format(date)
    .split("/")
    .reverse()
    .join("-");
};

/** result 31 Desember 2024 */
export const formatDate = (date: string | number) =>
  new Date(date).toLocaleDateString("id-ID", {
    year: "numeric",
    month: "long",
    day: "numeric",
  });

// Output: "10:00"
export const getHourAndMinute = (datetime: string): string => {
  const date = new Date(datetime);
  const hours = date.getUTCHours().toString().padStart(2, "0");
  const minutes = date.getUTCMinutes().toString().padStart(2, "0");
  return `${hours === "NaN" ? "-" : hours}:${
    minutes === "NaN" ? "-" : minutes
  }`;
};

/** result: "10:00:00" */
export const getTimes = (data: string | number) => {
  const date = new Date(data);

  // Gunakan Intl.DateTimeFormat biar pasti pakai zona waktu Asia/Jakarta
  const formatter = new Intl.DateTimeFormat("en-GB", {
    timeZone: "Asia/Jakarta",
    hour: "2-digit",
    minute: "2-digit",
    second: "2-digit",
    hour12: false, // biar 24 jam
  });

  // Formatnya jadi misalnya: "12:00:00"
  const timeString = formatter.format(date);
  return timeString;
};
