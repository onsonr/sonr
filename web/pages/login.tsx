import { Box, Center, Flex, Spacer, Tag, Text } from "@chakra-ui/react";
import {
  AppShell,
  Button,
  ButtonGroup,
  Card,
  CardBody,
  CardFooter,
  Form,
  Field,
  FormLayout,
  FormStep,
  FormStepper,
  FormValue,
  Loader,
  NextButton,
  PrevButton,
  Property,
  PropertyList,
  StepForm,
  PasswordInput,
  StepperCompleted,
  useModals,
  useSnackbar,
} from "@saas-ui/react";
import Link from "next/link";
import { useState } from "react";
export default function Login() {
  return (
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
  );
}
