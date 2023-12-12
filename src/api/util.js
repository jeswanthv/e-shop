// api.js
import axios from "axios";

const BASE_URL =
  "https://dfc5-2601-646-9980-1d80-e90e-ec90-ef52-5637.ngrok-free.app/";

// Function to fetch data from the API
export const fetchData = async (keyword) => {
  try {
    const response = await axios.get(`${BASE_URL}/search?keyword=${keyword}`);
    return response.data;
  } catch (error) {
    console.error("Error fetching data:", error);
    throw error;
  }
};

// Function to post data to the API
// export const postData = async (data) => {
//   try {
//     const response = await axios.post(`${BASE_URL}/data`, data);
//     return response.data;
//   } catch (error) {
//     console.error('Error posting data:', error);
//     throw error;
//   }
// };
