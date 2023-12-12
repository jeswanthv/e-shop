import { Search2Icon } from "@chakra-ui/icons";
import {
  Flex,
  Grid,
  Input,
  InputGroup,
  InputLeftElement,
  Spinner,
  Text,
} from "@chakra-ui/react";
import React, { useState } from "react";
import { fetchData } from "../api/util";
import Card from "../components/Card";
import Div from "../components/Div";
import Navbar from "../components/Navbar";

const Home = () => {
  const [data, setData] = useState([]);
  const [newData, setNewData] = useState("");
  const [loading, setLoading] = useState(false);

  const handleSearch = async (e) => {
    setLoading(true);
    try {
      const result = await fetchData(data);
      setNewData(result);
      console.log(result);
    } catch (error) {
      // Handle error
      console.log(error);
    }
    setLoading(false);
  };
  const handleKeyPress = (event) => {
    // Check if the pressed key is "Enter" (key code 13)
    if (event.key === "Enter") {
      // Perform the desired action, for example, submitting a form or calling a function
      console.log("Enter key pressed! Value:", data);
      handleSearch(data);
    }
  };
  return (
    <>
      <Navbar />
      {/* <Carousel /> */}
      <Div>
        <InputGroup mt={"100px"} w={"50%"} size="lg">
          <InputLeftElement pointerEvents="none">
            <Search2Icon color="gray.300" />
          </InputLeftElement>
          <Input
            type="text"
            value={data}
            onChange={(e) => setData(e.target.value)}
            onKeyDown={handleKeyPress}
            placeholder="Search"
          />
        </InputGroup>
      </Div>
      <Div>
        {loading ? (
          <Flex
            justifyContent={"center"}
            alignItems={"center"}
            direction={"column"}
          >
            <Spinner
              mt={"100px"}
              thickness="4px"
              speed="0.65s"
              emptyColor="gray.200"
              color="blue.500"
              size="xl"
            />
            <Text mt={"20px"}>
              Please wait until we fetch your best offers...
            </Text>
          </Flex>
        ) : (
          <Grid templateColumns="repeat(3, 1fr)" gap={6}>
            {newData &&
              newData.map((item, key) => <Card key={key} item={item} />)}
          </Grid>
        )}
      </Div>
    </>
  );
};

export default Home;
