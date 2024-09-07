import "./module.scss";

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

// Const
import routerConfig from "@/const/config/router";
import apiConfig from "@/const/config/api";

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

const barOption = {
  maintainAspectRatio: false,
  scales: {
    y: {
      beginAtZero: true,
      display: false,
    },
    x: {
      beginAtZero: true,
      display: false,
    },
    yAxes: {
      ticks: {
        stepSize: 1,
      },
    },
  },
};

const Dashboard = () => {
  // State
  const [pieChartData, setPieChartData] = useState(null);
  const [barChartData, setBarChartData] = useState(null);

  // Method
  const handleInitialization = useCallback(async () => {
    let response = null;
    response = await apiHandler.get(
      apiConfig.resource.STATISTIC_MOST_PUBLISHERS
    );

    // Make pie chart data
    const updatedPieChartData = {
      labels: [],
      datasets: [
        {
          data: [],
          backgroundColor: [],
        },
      ],
    };

    response.data.data.forEach((item) => {
      updatedPieChartData.labels.push(item.memberName);
      updatedPieChartData.datasets[0].data.push(item.numberOfPost);
      updatedPieChartData.datasets[0].backgroundColor.push(
        "#" + Math.floor(Math.random() * 16777215).toString(16)
      );
    });
    setPieChartData(updatedPieChartData);

    response = await apiHandler.get(apiConfig.resource.STATISTIC_MOST_COMMENTS);

    // Make most comments data
    const updatedBarCharData = {
      labels: [],
      datasets: [
        {
          label: "Number of comments",
          data: [],
          backgroundColor: [],
        },
      ],
    };

    response.data.data.forEach((item) => {
      updatedBarCharData.labels.push(item.documentName);
      updatedBarCharData.datasets[0].data.push(item.numberOfComment);
      updatedBarCharData.datasets[0].backgroundColor.push(
        "#" + Math.floor(Math.random() * 16777215).toString(16)
      );
    });
    setBarChartData(updatedBarCharData);
  }, []);

  // Side effect
  useEffect(() => {
    handleInitialization();
  }, [handleInitialization]);

  if (!pieChartData || !barChartData) {
    return null;
  }

  return (
    <>
      <div className='breadcrumb-container'>
        <Link to={routerConfig.routes.DASHBOARD} className='breadcrumb--item'>
          <span className='breadcrumb--item--inner'>
            <span className='breadcrumb--item-title'>Dashboard</span>
          </span>
        </Link>
      </div>

      <div className='section two-columns'>
        <div>
          <div className='type-title'>The most publishers</div>
          <div className='chart'>
            <Pie data={pieChartData} height='200px' />
          </div>
        </div>
        <div>
          <div className='type-title'>The most comments</div>
          <div className='chart'>
            <Bar data={barChartData} options={barOption} />
          </div>
        </div>
      </div>
    </>
  );
};

export default Dashboard;
