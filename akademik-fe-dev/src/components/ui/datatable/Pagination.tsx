"use client";

import ReactPaginate from "react-paginate";

import styles from "@/styles/pagination.module.css";

interface PaginationProps {
  pageCount: number;
  pageOffset: number;
  onPageChange: (e: { selected: number }) => void;
}

export const Pagination: React.FC<PaginationProps> = ({
  pageCount,
  pageOffset,
  onPageChange,
}) => {
  return (
    <ReactPaginate
      breakLabel={<span className={styles.break_label}>...</span>}
      nextLabel={<span className={styles.next_prev_label}>Selanjutnya</span>}
      previousLabel={<span className={styles.next_prev_label}>Sebelumnya</span>}
      onPageChange={onPageChange}
      pageCount={pageCount}
      renderOnZeroPageCount={null}
      containerClassName={styles.container_pagination}
      pageClassName={styles.page}
      activeClassName={styles.active_page}
      forcePage={pageOffset - 1}
      marginPagesDisplayed={3}
      disableInitialCallback={true}
      nextClassName={styles.next_prev}
      previousClassName={styles.next_prev}
      disabledClassName={styles.disabled}
    />
  );
};
