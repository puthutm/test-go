"use client";

import React, { useState, useEffect } from "react";
import { useForm, Controller } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import {
  Card,
  CardBody,
  CardHeader,
  Col,
  Row,
  Label,
  Button,
} from "reactstrap";
import {
  PieChart,
  Pie,
  Cell,
  ResponsiveContainer,
  Legend,
  Tooltip,
} from "recharts";

import {
  AssessmentWeightSchema,
  AssessmentWeightFormType,
} from "@/lib/validations/settings/assessment-weight/form-assessment-weight";
import { updateAssessmentWeight } from "@/services/api/settings/assessment-weight/update-assessment-weight";
import { getAssessmentWeight } from "@/services/api/settings/assessment-weight/get-assessment-weight";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";

const AssessmentWeightPage = () => {
  const [isLoading, setIsLoading] = useState(false);
  const { setModalConfirmationState } = useModalConfirmationContext();

  const {
    control,
    handleSubmit,
    setValue,
    watch,
    formState: { errors, isSubmitting },
  } = useForm<AssessmentWeightFormType>({
    resolver: zodResolver(AssessmentWeightSchema),
    defaultValues: {
      attitude_behavior_percentage: "25",
      task_percentage: "25",
      uts_percentage: "25",
      uas_percentage: "25",
    },
    mode: "onChange",
  });

  // Watch all values for dynamic chart
  const watchedValues = watch();
  const totalPercentage =
    Number(watchedValues.attitude_behavior_percentage || 0) +
    Number(watchedValues.task_percentage || 0) +
    Number(watchedValues.uts_percentage || 0) +
    Number(watchedValues.uas_percentage || 0);

  // Chart data
  const chartData = [
    {
      name: "Sikap/Perilaku",
      value: Number(watchedValues.attitude_behavior_percentage || 0),
      color: "#405189",
    },
    {
      name: "Tugas",
      value: Number(watchedValues.task_percentage || 0),
      color: "#0ab39c",
    },
    {
      name: "UTS",
      value: Number(watchedValues.uts_percentage || 0),
      color: "#f7b84b",
    },
    {
      name: "UAS",
      value: Number(watchedValues.uas_percentage || 0),
      color: "#f06548",
    },
  ];

  // Load existing data
  useEffect(() => {
    const loadData = async () => {
      try {
        setIsLoading(true);
        const response = await getAssessmentWeight();

        if (response?.data) {
          setValue(
            "attitude_behavior_percentage",
            String(response.data.attitude_behavior_percentage)
          );
          setValue("task_percentage", String(response.data.task_percentage));
          setValue("uts_percentage", String(response.data.uts_percentage));
          setValue("uas_percentage", String(response.data.uas_percentage));
        }
      } catch (error) {
        console.error("Error loading assessment weight data:", error);
      } finally {
        setIsLoading(false);
      }
    };

    loadData();
  }, [setValue]);

  const onSubmit = async (data: AssessmentWeightFormType) => {
    try {
      const response = await updateAssessmentWeight(data);
      if (response.error) {
        throw new Error(response.message);
      }

      setModalConfirmationState({
        open: true,
        message: "Berhasil mengupdate bobot penilaian",
        state: "success",
      });
    } catch (error: any) {
      setModalConfirmationState({
        open: true,
        message: error.message || "Terjadi kesalahan saat mengupdate data",
        state: "failed",
      });
    }
  };

  return (
    <Row>
      <Col lg={8}>
        <Card>
          <CardHeader>
            <h4 className="card-title mb-0">Pengaturan Bobot Penilaian</h4>
          </CardHeader>
          <CardBody>
            <form onSubmit={handleSubmit(onSubmit)}>
              <Row className="gy-4">
                {/* Sikap/Perilaku */}
                <Col lg={6}>
                  <Label className="form-label">
                    Persentase Sikap/Perilaku:{" "}
                    {watchedValues.attitude_behavior_percentage}%
                  </Label>
                  <Controller
                    name="attitude_behavior_percentage"
                    control={control}
                    render={({ field }) => (
                      <input
                        type="range"
                        className="form-range"
                        min="0"
                        max="100"
                        step="0.25"
                        value={field.value}
                        onChange={(e) => field.onChange(e.target.value)}
                        disabled={isLoading}
                      />
                    )}
                  />
                  {errors.attitude_behavior_percentage && (
                    <div className="text-danger fs-6 mt-1">
                      {errors.attitude_behavior_percentage.message}
                    </div>
                  )}
                </Col>

                {/* Tugas */}
                <Col lg={6}>
                  <Label className="form-label">
                    Persentase Tugas: {watchedValues.task_percentage}%
                  </Label>
                  <Controller
                    name="task_percentage"
                    control={control}
                    render={({ field }) => (
                      <input
                        type="range"
                        className="form-range"
                        min="0"
                        max="100"
                        step="0.25"
                        value={field.value}
                        onChange={(e) => field.onChange(e.target.value)}
                        disabled={isLoading}
                        data-rangeslider
                        data-slider-color="primary"
                      />
                    )}
                  />
                  {errors.task_percentage && (
                    <div className="text-danger fs-6 mt-1">
                      {errors.task_percentage.message}
                    </div>
                  )}
                </Col>

                {/* UTS */}
                <Col lg={6}>
                  <Label className="form-label">
                    Persentase UTS: {watchedValues.uts_percentage}%
                  </Label>
                  <Controller
                    name="uts_percentage"
                    control={control}
                    render={({ field }) => (
                      <input
                        type="range"
                        className="form-range"
                        min="0"
                        max="100"
                        step="0.25"
                        value={field.value}
                        onChange={(e) => field.onChange(e.target.value)}
                        disabled={isLoading}
                      />
                    )}
                  />
                  {errors.uts_percentage && (
                    <div className="text-danger fs-6 mt-1">
                      {errors.uts_percentage.message}
                    </div>
                  )}
                </Col>

                {/* UAS */}
                <Col lg={6}>
                  <Label className="form-label">
                    Persentase UAS: {watchedValues.uas_percentage}%
                  </Label>
                  <Controller
                    name="uas_percentage"
                    control={control}
                    render={({ field }) => (
                      <input
                        type="range"
                        className="form-range"
                        min="0"
                        max="100"
                        step="0.25"
                        value={field.value}
                        onChange={(e) => field.onChange(e.target.value)}
                        disabled={isLoading}
                      />
                    )}
                  />
                  {errors.uas_percentage && (
                    <div className="text-danger fs-6 mt-1">
                      {errors.uas_percentage.message}
                    </div>
                  )}
                </Col>
              </Row>

              {/* Total Indicator */}
              <Row className="mt-4">
                <Col lg={12}>
                  <div
                    className={`alert ${
                      totalPercentage === 100
                        ? "alert-success"
                        : "alert-warning"
                    } text-center`}
                  >
                    <h5 className="mb-0">
                      Total: {totalPercentage}%
                      {totalPercentage === 100 ? (
                        <i className="ri-check-line ms-2"></i>
                      ) : (
                        <i className="ri-error-warning-line ms-2"></i>
                      )}
                    </h5>
                    {totalPercentage !== 100 && (
                      <small>Total persentase harus sama dengan 100%</small>
                    )}
                  </div>
                </Col>
              </Row>

              {/* Submit Button */}
              <Row className="mt-3">
                <Col lg={12} className="text-end">
                  <Button
                    type="submit"
                    color="primary"
                    disabled={
                      isSubmitting || isLoading || totalPercentage !== 100
                    }
                    className="btn"
                  >
                    Simpan Pengaturan
                  </Button>
                </Col>
              </Row>
            </form>
          </CardBody>
        </Card>
      </Col>

      {/* Chart Section */}
      <Col lg={4}>
        <Card>
          <CardHeader>
            <h4 className="card-title mb-0">Visualisasi Distribusi Bobot</h4>
          </CardHeader>
          <CardBody>
            <div style={{ width: "100%", height: "300px" }}>
              <ResponsiveContainer>
                <PieChart>
                  <Pie
                    data={chartData}
                    cx="50%"
                    cy="50%"
                    innerRadius={60}
                    outerRadius={120}
                    paddingAngle={2}
                    dataKey="value"
                  >
                    {chartData.map((entry, index) => (
                      <Cell key={`cell-${index}`} fill={entry.color} />
                    ))}
                  </Pie>
                  <Tooltip
                    formatter={(value: any, name: any) => [`${value}%`, name]}
                  />
                  <Legend />
                </PieChart>
              </ResponsiveContainer>
            </div>

            {/* Summary Stats */}
            <div className="mt-3">
              <div className="d-flex justify-content-between border-bottom pb-2 mb-2">
                <span>Total Distribusi:</span>
                <span
                  className={`fw-bold ${
                    totalPercentage === 100 ? "text-success" : "text-warning"
                  }`}
                >
                  {totalPercentage}%
                </span>
              </div>
              <div className="text-muted small">
                {totalPercentage === 100
                  ? "✓ Distribusi sudah seimbang"
                  : `⚠ Masih kurang ${100 - totalPercentage}%`}
              </div>
            </div>
          </CardBody>
        </Card>
      </Col>
    </Row>
  );
};

export default AssessmentWeightPage;
