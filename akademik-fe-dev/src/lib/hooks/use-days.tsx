export const useWeekDaysOptions = (): OptionType[] => {
  const days = [
    { id: "Minggu", en: "sunday" },
    { id: "Senin", en: "monday" },
    { id: "Selasa", en: "tuesday" },
    { id: "Rabu", en: "wednesday" },
    { id: "Kamis", en: "thursday" },
    { id: "Jumat", en: "friday" },
    { id: "Sabtu", en: "saturday" },
  ];

  return days.map((day) => ({
    label: day.id,
    value: day.en,
  }));
};
