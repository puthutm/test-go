"use client";

import {  useState } from "react";
import { signOut } from "next-auth/react";
import { useSearchParams, usePathname, useRouter } from "next/navigation";
import { useCreditLimitColumns } from "./columns";
import { SearchIcon } from "@/components/icons/search";
import {
  Button,
  Card,
  CardBody,
  CardHeader,
  Col,
  Input,
  Row,
} from "reactstrap";
import DataTables from "@/components/ui/datatable";


import { useModalContext } from "@/lib/hooks/use-modal";
import ModalCreditLimit from "./ModalCreditLimit";

import { deleteSksLimit } from "@/services/api/settings/sks-limit/delete-sks-limits";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { ModalDeleteConfirmation } from "@/components/ui/modal-delete-confirmation";
import { useDebouncedCallback } from "use-debounce";

import { AddIcon } from "@/components/icons/add";
import { DeleteIcon } from "@/components/icons/delete";
import Link from "next/link";

const CreditLimitClientPage = ({
  dataSksLimit,
}: {
  dataSksLimit: ApiResponse<PaginationData<ISksLimit>>;
}) => {
  const router = useRouter();
  const pathname = usePathname();
  const searchParams = useSearchParams();
  const params = new URLSearchParams(searchParams);

  const { columns } = useCreditLimitColumns();

  const [isLoadingDelete, setIsLoadingDelete] = useState<boolean>(false);
  const { setModalState } = useModalContext();
  const { modalConfirmationState, setModalConfirmationState } =
    useModalConfirmationContext();


  //! handle search
  const handleSearch = useDebouncedCallback((search: string) => {
    if (search.trim().length > 0) {
      params.set("search", search);
    } else {
      params.delete("search");
    }
    params.set("page", "1"); // Reset to first page on search
    router.push(`${pathname}?${params.toString()}`);
  }, 1000);


    //! event handle page pagination
  const handlePagination = (newPage: number) => {
    if (newPage) {
      params.set("page", (newPage + 1).toString());
    } else {
      params.delete("page");
    }

    router.push(`${pathname}?${params.toString()}`);
  };

    const handleDelete = async (sksLimitsId: string) => {
      setIsLoadingDelete(true);
      try {
        const response = await deleteSksLimit(sksLimitsId);
        if (response.error) {
        setModalConfirmationState((prev) => ({
          ...prev,
          id: null,
          open:true,
          state: "failed",
          message: response.message || "Gagal menghapus data",
        }));
        return
        }
  
        setModalConfirmationState((prev) => ({
          ...prev,
          open:true,
          id: null,
          state: "success",
          message: "Data berhasil dihapus",
        }));
      } catch (err: any) {
        throw new Error(err.message);
      } finally {
        setIsLoadingDelete(false);
      }
    };

      if (dataSksLimit.status === 401) {
    signOut();
  }

  return (
    <Row>
      <Col>
        <ModalCreditLimit />
        <ModalDeleteConfirmation
          isLoading={isLoadingDelete}
          onDelete={async () => {
            await handleDelete(modalConfirmationState.id as string);
          }}
        />
        <Card className="p-0">
          <CardHeader>
            <div className="gap-2 d-flex align-items-center justify-content-between w-100">
              <div className="d-flex gap-2">
                <div className="form-icon">
                  <Input
                    type="text"
                    className="form-control form-control-icon"
                    id="iconInput"
                    placeholder="Cari"
                    onChange={(event) =>
                      handleSearch((event.target as HTMLInputElement).value)
                    }
                  />
                  <i><SearchIcon /></i>
                </div>
              </div>

              <div className="d-flex gap-2">
                {/* Trash */}
                <Link
                  href={"/settings/credit-limit/trash"}
                  className="btn btn-danger"
                >
                 <DeleteIcon color="#fff"/>
                  Trash
                </Link>
                {/*//! button tambah */}
                <Button
                  className="btn  d-flex align-items-center gap-2"
                  style={{
                    color: "#10487A",
                    border: "1px solid #10487A",
                  }}
                  color="transparent"
                  onClick={() => {
                      setModalState((prev) => ({
                        ...prev,
                        open: !prev.open,
                        state:'add',
                        id:null,
                      }));
                  }}
                >
                  <AddIcon color="#10487A" />
                  Tambah
                </Button>
              </div>
            </div>
          </CardHeader>

          <CardBody>
            <Col className="table-responsive" sm={12}>
              <DataTables
                columns={columns}
                data={dataSksLimit?.data}
                pageCount={dataSksLimit?.data?.metadata.total_page as number}
                pagination={dataSksLimit?.data?.metadata}
                setPagination={handlePagination}
                total={dataSksLimit?.data?.metadata.total_data as number}
              />
            </Col>
          </CardBody>
        </Card>
      </Col>
    </Row>
  );
};

export default CreditLimitClientPage;
