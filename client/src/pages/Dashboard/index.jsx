import "./module.css";

import { useState, useEffect, useCallback } from "react";
import { Link } from "react-router-dom";
import { Pie, Bar } from "react-chartjs-2";
import {
  Chart as ChartJS,
  ArcElement,
  CategoryScale,
  LinearScale,
  BarController,
  BarElement,
  Title,
  Tooltip,
  Legend,
} from "chart.js";
import { cloneDeep } from "lodash";

// Const
import apiConfig from "@/const/config/api";

// Component
import Breadcrumbs from "@/components/Breadcrumbs";
import MainTitle from "@/components/MainTitle";
import Form from "@/components/Form";
import Spacer from "@/components/Spacer";

// Util
import apiHandler from "@/util/api.util";

ChartJS.register(
  ArcElement,
  CategoryScale,
  LinearScale,
  BarController,
  BarElement,
  Title,
  Tooltip,
  Legend
);

const defaultPieChart = {
  labels: [],
  datasets: [
    {
      data: [],
      backgroundColor: [],
    },
  ],
};

const defaultBarChart = {
  labels: [],
  datasets: [
    {
      borderWidth: 1,
      data: [],
      formattedValue: [],
      backgroundColor: [],
    },
  ],
};

const pieOption = {
  maintainAspectRatio: false,
  plugins: {
    legend: {
      display: true,
      position: "right",
    },
  },
};

const barOption = {
  maintainAspectRatio: false,
  plugins: {
    legend: {
      display: false,
    },
    tooltip: {
      callbacks: {
        label: function (data) {
          return data.dataset.formattedValue[data.dataIndex];
        },
      },
    },
  },
  y: {
    ticks: {
      stepSize: 1,
      suggestedMin: "min-int-value",
      suggestedMax: "max-int-value",
    },
  },
};

const Dashboard = () => {
  const linkList = [{ to: "/", label: "Dashboard" }];

  // State
  const [mostPublisherPieChart, setMostPublisherPieChart] = useState(null);
  const [mostPublisherBarChart, setMostPublisherBarChart] = useState(null);
  const [mostCommentBarChart, setMostCommentBarChart] = useState(null);

  // Method
  const handleInitialization = useCallback(async () => {
    const mostPublisherResponse = await apiHandler.get(
      apiConfig.resource.STATISTIC_MOST_PUBLISHERS
    );

    const mostCommentResponse = await apiHandler.get(
      apiConfig.resource.STATISTIC_MOST_COMMENTS
    );

    // Make pie chart data
    const updatedMostPublisherPieChart = cloneDeep(defaultPieChart);
    mostPublisherResponse.data.data?.forEach((item) => {
      updatedMostPublisherPieChart.labels.push(item.memberName);
      updatedMostPublisherPieChart.datasets[0].data.push(item.numberOfPost);
      updatedMostPublisherPieChart.datasets[0].backgroundColor.push(
        "#" + Math.floor(Math.random() * 16777215).toString(16)
      );
    });
    setMostPublisherPieChart(updatedMostPublisherPieChart);

    // Make bar chart data
    const updatedMostPublishBarChart = cloneDeep(defaultBarChart);
    mostPublisherResponse.data.data?.forEach((item) => {
      updatedMostPublishBarChart.labels.push(item.memberName);
      updatedMostPublishBarChart.datasets[0].data.push(item.numberOfPost);
      updatedMostPublishBarChart.datasets[0].backgroundColor.push(
        "#" + Math.floor(Math.random() * 16777215).toString(16)
      );
    });
    setMostPublisherBarChart(updatedMostPublishBarChart);

    const updatedMostCommentBarChart = cloneDeep(defaultBarChart);
    mostCommentResponse.data.data?.forEach((item) => {
      updatedMostCommentBarChart.labels.push(item.categoryName);
      updatedMostCommentBarChart.datasets[0].formattedValue.push(
        item.documentName
      );
      updatedMostCommentBarChart.datasets[0].data.push(item.numberOfComment);
      updatedMostCommentBarChart.datasets[0].backgroundColor.push(
        "#" + Math.floor(Math.random() * 16777215).toString(16)
      );
    });
    setMostCommentBarChart(updatedMostCommentBarChart);
  }, []);

  // Side effect
  useEffect(() => {
    handleInitialization();

    const intervalId = setInterval(() => {
      handleInitialization();
    }, 1000 * 60 * 60);

    return () => {
      clearInterval(intervalId);
    };
  }, [handleInitialization]);

  if (
    !mostPublisherPieChart ||
    !mostPublisherBarChart ||
    !mostCommentBarChart
  ) {
    return null;
  }

  return (
    <>
      <Breadcrumbs linkList={linkList} />
      <div className='custom-container primary-bg'>
        <MainTitle>Top 10 Publishers</MainTitle>
        <Form>
          <div className='publisher-container'>
            <div className='chart'>
              {mostPublisherPieChart.labels.length === 0 ? (
                <MainTitle extraClasses={["no-data"]}>No data</MainTitle>
              ) : (
                <Pie data={mostPublisherPieChart} options={pieOption} />
              )}
            </div>
            <div className='chart'>
              {mostPublisherBarChart.labels.length === 0 ? (
                <MainTitle extraClasses={["no-data"]}>No data</MainTitle>
              ) : (
                <Bar data={mostPublisherBarChart} options={barOption} />
              )}
            </div>
          </div>
        </Form>
        <Spacer extraClasses={["mt-4"]} />
        <MainTitle>Top 10 Comments</MainTitle>
        <Form>
          <div className='chart'>
            {mostCommentBarChart.labels.length === 0 ? (
              <MainTitle extraClasses={["no-data"]}>No data</MainTitle>
            ) : (
              <Bar data={mostCommentBarChart} options={barOption} />
            )}
          </div>
        </Form>
      </div>
    </>
  );
};

export default Dashboard;
