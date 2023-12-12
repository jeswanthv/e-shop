import { Flex } from "@chakra-ui/react";
import React from "react";

const Div = ({ children }) => {
  return (
    <Flex
      margin={"auto"}
      maxW={"1200px"}
      justifyContent={"center"}
      alignItems={"center"}
    >
      {children}
    </Flex>
  );
};

export default Div;
