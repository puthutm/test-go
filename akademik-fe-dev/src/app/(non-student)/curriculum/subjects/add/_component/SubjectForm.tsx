"use client";

import { classMerge } from "@/lib/utils/class-merge";
import {
  SubjectFormSchema,
  SubjectFormType,
  SubjectInitialValue,
} from "@/lib/validations/curriculum/subject";
// import { useCreateSubject } from "@/services/api/curriculum/subjects/create";
import { useGetSearchCourseGroups } from "@/services/api/reference/course-groups/get-search";
import { useGetSearchCourseTypes } from "@/services/api/reference/course-types/get-search";
import { useGetSearchStudyProgram } from "@/services/api/reference/study-program/get-search";
import { zodResolver } from "@hookform/resolvers/zod";
import React, { useEffect, useState } from "react";
import { Controller, useForm } from "react-hook-form";
import { Col, Input, Label, Row } from "reactstrap";

const AddSubjectForm = () => {
  const [queryParams] = useState<QueryParam>({
    page: 1,
    // filter: null,
    // page_size: 1000,
    sort_by: null,
    sort_direction: null,
  });

  // const { mutateAsync: onCreate, isSuccess: isCreateSuccess } =
  //   useCreateSubject();

  const {
    data: studyProgramData,
    isLoading: isStudyProgramLoading,
    refetch: refetchStudyProgram,
  } = useGetSearchStudyProgram(queryParams);

  const {
    data: courseTypesData,
    // isLoading: isCourseTypesLoading,
    refetch: refetchCourseTypes,
  } = useGetSearchCourseTypes(queryParams);

  const {
    data: courseGroupsData,
    // isLoading: isCourseGroupsLoading,
    refetch: refetchCourseGroups,
  } = useGetSearchCourseGroups(queryParams);

  const {
    formState: { errors,  },
    handleSubmit,
    control,
    // setValue,
    // reset,
  } = useForm<SubjectFormType>({
    resolver: zodResolver(SubjectFormSchema),
    defaultValues: SubjectInitialValue,
  });

  const onSubmit = async (data: SubjectFormType) => {
    console.log("ini data", data);
  };

  useEffect(() => {
    refetchStudyProgram();
    refetchCourseTypes();
    refetchCourseGroups();
  });

  return (
    <form onSubmit={handleSubmit(onSubmit)} autoComplete="off">
      <Row>
        <Col className="d-flex flex-column gap-3">
          <Row>
            <Label
              htmlFor="curriculum_year_id"
              className="form-label required mb-1"
            >
              Tahun Kurikulum
            </Label>
            <Controller
              control={control}
              name="curriculum_year_id"
              render={({ field }) => (
                <div>
                  <select
                    {...field}
                    className={classMerge(
                      errors.curriculum_year_id ? "is-invalid" : "",
                      "form-select"
                    )}
                    id="curriculum_year_id"
                    disabled={isStudyProgramLoading}
                  >
                    <option value="">Pilih Tahun Kurikulum</option>
                    {studyProgramData?.data?.map((item: any) => (
                      <option key={item.id} value={item.id}>
                        {item.name}
                      </option>
                    ))}
                  </select>
                </div>
              )}
            />
            {errors?.curriculum_year_id && (
              <div className="text-danger">
                {errors?.curriculum_year_id.message}
              </div>
            )}
          </Row>

          {/* Study Program */}
          <Row>
            <Label
              htmlFor="study_program_id"
              className="form-label mb-1 required"
            >
              Program Studi
            </Label>
            <Controller
              control={control}
              name="study_program_id"
              render={({ field }) => (
                <div>
                  <select
                    {...field}
                    className={classMerge(
                      errors.study_program_id ? "is-invalid" : "",
                      "form-select form-control-icon"
                    )}
                    id="study_program_id"
                    disabled={isStudyProgramLoading}
                  >
                    <option value="">Pilih Program Studi</option>
                    {studyProgramData?.data?.map((item: any) => (
                      <option key={item.id} value={item.id}>
                        {item.name}
                      </option>
                    ))}
                  </select>
                </div>
              )}
            />
            {errors?.study_program_id && (
              <div className="text-danger">
                {errors?.study_program_id.message}
              </div>
            )}
          </Row>

          {/* code */}
          <Row>
            <Label htmlFor="code" className="form-label mb-1 required">
              Kode
            </Label>
            <Controller
              control={control}
              name="code"
              render={({ field }) => (
                <div>
                  <Input
                    {...field}
                    type="text"
                    id="code"
                    placeholder="Masukkan Kode"
                    invalid={!!errors.code}
                  />
                </div>
              )}
            />
            {errors?.code && (
              <div className="text-danger">{errors?.code.message}</div>
            )}
          </Row>

          {/* name_id */}
          <Row>
            <Label htmlFor="name_id" className="form-label mb-1 required">
              Nama Mata Kuliah (Bahasa Indonesia)
            </Label>
            <Controller
              control={control}
              name="name_id"
              render={({ field }) => (
                <div>
                  <Input
                    {...field}
                    type="text"
                    id="name"
                    placeholder="Masukkan Nama"
                    invalid={!!errors.name_id}
                  />
                </div>
              )}
            />
            {errors?.name_id && (
              <div className="text-danger">{errors?.name_id.message}</div>
            )}
          </Row>

          {/* name_en */}
          <Row>
            <Label htmlFor="name_en" className="form-label mb-1 required">
              Nama Mata Kuliah (Bahasa Inggris)
            </Label>
            <Controller
              control={control}
              name="name_en"
              render={({ field }) => (
                <div>
                  <Input
                    {...field}
                    type="text"
                    id="name"
                    placeholder="Masukkan Nama"
                    invalid={!!errors.name_en}
                  />
                </div>
              )}
            />
            {errors?.name_en && (
              <div className="text-danger">{errors?.name_en.message}</div>
            )}
          </Row>

          {/* course_type_id */}
          <Row>
            <Label
              htmlFor="course_type_id"
              className="form-label mb-1 required"
            >
              Jenis Mata Kuliah
            </Label>
            <Controller
              control={control}
              name="course_type_id"
              render={({ field }) => (
                <div>
                  <select
                    {...field}
                    className={classMerge(
                      errors.course_type_id ? "is-invalid" : "",
                      "form-select"
                    )}
                    id="course_type_id"
                  >
                    <option value="">Pilih Jenis Mata Kuliah</option>
                    {courseTypesData?.data?.map((item: any) => (
                      <option key={item.id} value={item.id}>
                        {item.name}
                      </option>
                    ))}
                  </select>
                </div>
              )}
            />
            {errors?.course_type_id && (
              <div className="text-danger">
                {errors?.course_type_id.message}
              </div>
            )}
          </Row>

          {/* course_group_id */}
          <Row>
            <Label
              htmlFor="course_group_id"
              className="form-label mb-1 required"
            >
              Kelompok Mata Kuliah
            </Label>
            <Controller
              control={control}
              name="course_group_id"
              render={({ field }) => (
                <div>
                  <select
                    {...field}
                    className={classMerge(
                      errors.course_group_id ? "is-invalid" : "",
                      "form-select"
                    )}
                    id="course_group_id"
                  >
                    <option value="">Pilih Kelompok Mata Kuliah</option>
                    {courseGroupsData?.data?.map((item: any) => (
                      <option key={item.id} value={item.id}>
                        {item.name}
                      </option>
                    ))}
                  </select>
                </div>
              )}
            />
            {errors?.course_group_id && (
              <div className="text-danger">
                {errors?.course_group_id.message}
              </div>
            )}
          </Row>

          {/* face_to_face_sks */}
          <Row>
            <Label
              htmlFor="face_to_face_sks"
              className="form-label mb-1 required"
            >
              SKS Tatap Muka
            </Label>
            <Controller
              control={control}
              name="face_to_face_sks"
              render={({ field }) => (
                <div>
                  <Input
                    {...field}
                    type="number"
                    id="face_to_face_sks"
                    placeholder="Masukkan SKS"
                    invalid={!!errors.face_to_face_sks}
                  />
                </div>
              )}
            />
            {errors?.face_to_face_sks && (
              <div className="text-danger">
                {errors?.face_to_face_sks.message}
              </div>
            )}
          </Row>

          {/* practicum_sks */}
          <Row>
            <Label htmlFor="practicum_sks" className="form-label mb-1 required">
              SKS Praktikum
            </Label>
            <Controller
              control={control}
              name="practicum_sks"
              render={({ field }) => (
                <div>
                  <Input
                    {...field}
                    type="number"
                    id="practicum_sks"
                    placeholder="Masukkan SKS"
                    invalid={!!errors.practicum_sks}
                  />
                </div>
              )}
            />
            {errors?.practicum_sks && (
              <div className="text-danger">{errors?.practicum_sks.message}</div>
            )}
          </Row>

          {/* field_practice_sks */}
          <Row>
            <Label
              htmlFor="field_practice_sks"
              className="form-label mb-1 required"
            >
              SKS Praktik Lapangan
            </Label>
            <Controller
              control={control}
              name="field_practice_sks"
              render={({ field }) => (
                <div>
                  <Input
                    {...field}
                    type="number"
                    id="field_practice_sks"
                    placeholder="Masukkan SKS"
                    invalid={!!errors.field_practice_sks}
                  />
                </div>
              )}
            />
            {errors?.field_practice_sks && (
              <div className="text-danger">
                {errors?.field_practice_sks.message}
              </div>
            )}
          </Row>

          {/* simulation_sks */}
          <Row>
            <Label
              htmlFor="simulation_sks"
              className="form-label mb-1 required"
            >
              SKS Simulasi
            </Label>
            <Controller
              control={control}
              name="simulation_sks"
              render={({ field }) => (
                <div>
                  <Input
                    {...field}
                    type="number"
                    id="simulation_sks"
                    placeholder="Masukkan SKS"
                    invalid={!!errors.simulation_sks}
                  />
                </div>
              )}
            />
            {errors?.simulation_sks && (
              <div className="text-danger">
                {errors?.simulation_sks.message}
              </div>
            )}
          </Row>

          {/* total_sks */}
          <Row>
            <Label htmlFor="total_sks" className="form-label mb-1 required">
              Total SKS
            </Label>
            <Controller
              control={control}
              name="total_sks"
              render={({ field }) => (
                <div>
                  <Input
                    {...field}
                    type="number"
                    id="total_sks"
                    placeholder="Masukkan SKS"
                    disabled
                    invalid={!!errors.total_sks}
                  />
                </div>
              )}
            />
            {errors?.total_sks && (
              <div className="text-danger">{errors?.total_sks.message}</div>
            )}
          </Row>
        </Col>
        <Col>
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Voluptate
          recusandae sequi adipisci reiciendis officia, sapiente provident omnis
          veniam cupiditate delectus tempora eveniet ipsa veritatis,
          perspiciatis eligendi. Voluptas saepe sit cumque animi! Autem,
          consectetur? Iusto assumenda ipsum dolor vero aliquid voluptates quo
          libero perferendis aut, quos quas veniam quis, animi delectus commodi
          dolore officiis non, saepe debitis voluptate! Aut blanditiis
          voluptatum earum deleniti! Reprehenderit, quas! Iusto voluptatum eaque
          recusandae vel quaerat! Illo exercitationem iure tempore assumenda
          quidem nam in, suscipit dolore omnis aspernatur impedit voluptates
          minima id dignissimos repudiandae magnam deserunt velit. Praesentium
          unde iste architecto libero modi dolore ipsa voluptatum!
        </Col>
      </Row>
    </form>
  );
};

export default AddSubjectForm;
