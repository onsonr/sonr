import { Box, Center, Flex, Spacer, Tag, Text } from "@chakra-ui/react";
import { AppShell, Button, Card, CardBody, Form, Field } from "@saas-ui/react";
import Head from "next/head";
import Link from "next/link";

export default function Login() {
  return (
    <>
      <Head>
        <title>Sonr Sandbox | Login</title>
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
            <Link href="/">
              <Button variant="unstyled">
                <Text fontSize="xl" fontWeight="bold" paddingTop={1}>
                  Sonr Sandbox
                </Text>
              </Button>
            </Link>
            <Spacer />
            <Box>
              <Tag size="lg">v0.1.12</Tag>
            </Box>
          </Flex>
        }
      >
        <Center>
          <Box
            as="main"
            height={{
              base: "100%", // 0-48em
              md: "50%", // 48em-80em,
              xl: "50%", // 80em+
            }}
            width={[
              "100%", // 0-30em
              "90%", // 30em-48em
              "75%", // 48em-62em
              "50%", // 62em+
            ]}
          >
            <Card
              title="Import Existing Sonr Account"
              padding={4}
              action={
                <Link href="/">
                  <Button label="Cancel" variant="ghost" />
                </Link>
              }
            >
              <CardBody>
                <Form onSubmit={() => {}}>
                  <Field type="text" name="address" label="Address" />
                </Form>
              </CardBody>
            </Card>
          </Box>
        </Center>
      </AppShell>
    </>
  );
}
