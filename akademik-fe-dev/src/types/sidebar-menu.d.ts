interface SidebarMenu {
  id: string;
  label: string;
  order: number;
  path: string;
  icon: string;
  children: [
    {
      id: string;
      label: string;
      order: number;
      path: string;
      icon: string;
      parent_id;
    }
  ];
}
