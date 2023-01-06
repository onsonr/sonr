import {
  Box,
  Center,
  Flex,
  Spacer,
  Tag,
  Text,
  useColorMode,
} from "@chakra-ui/react";
import { AppShell, Button, EmptyState } from "@saas-ui/react";
import Head from "next/head";
import Link from "next/link";

export default function Home() {
  const { colorMode, toggleColorMode } = useColorMode();
  if (colorMode === "light") {
    toggleColorMode();
  }
  return (
    <>
      <Head>
        <title>Sonr Sandbox</title>
        <meta
          name="description"
          content="API Test Utility for the Sonr Blockchain"
        />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <AppShell
        navbar={
          <Flex borderBottomWidth="1px" py="2" px="4">
            <Box>
              <Text fontSize="xl" fontWeight="bold">
                Sonr Sandbox
              </Text>
            </Box>
            <Spacer />
            <Box>
              <Tag size="lg">v0.1.0</Tag>
            </Box>
          </Flex>
        }
      >
        <Center height="100vh" marginTop="15vh">
          <EmptyState
            colorScheme="primary"
            // icon={FiUsers}
            title="No Accounts Yet"
            description="You haven't imported any customers yet."
            actions={
              <>
                <Link href="/login">
                  <Button variant="primary">Import account</Button>
                </Link>
                <Link href="/signup">
                  <Button>Create Account</Button>
                </Link>
              </>
            }
          />
        </Center>
      </AppShell>
    </>
  );
}
