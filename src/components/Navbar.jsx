import { MoonIcon, SunIcon } from "@chakra-ui/icons";
import {
  Box,
  Button,
  Flex,
  Stack,
  Text,
  useColorMode,
  useColorModeValue,
} from "@chakra-ui/react";

export default function Navbar() {
  const { colorMode, toggleColorMode } = useColorMode();
  return (
    <>
      <Box
        position={"fixed"}
        zIndex={1}
        w={"100%"}
        opacity={0.8}
        bg={useColorModeValue("gray.100", "gray.900")}
        px={4}
      >
        <Flex
          maxW={"1200px"}
          h={16}
          margin={"auto"}
          alignItems={"center"}
          justifyContent={"space-between"}
        >
          <Box>
            <Text fontSize={"3xl"}>Smart Offers</Text>
          </Box>

          <Flex alignItems={"center"}>
            <Stack direction={"row"} spacing={7}>
              <Button onClick={toggleColorMode}>
                {colorMode === "light" ? <MoonIcon /> : <SunIcon />}
              </Button>
            </Stack>
          </Flex>
        </Flex>
      </Box>
    </>
  );
}
