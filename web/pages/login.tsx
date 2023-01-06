import { Box, Divider, Flex, Spacer, Tag, Text } from "@chakra-ui/react";
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
export default function Login() {
  return (
    <AppShell
      navbar={
        <Flex borderBottomWidth="1px" py="2" px="4">
          <Link href="/">
            <Text fontSize="xl" fontWeight="bold" paddingTop={1}>
              Sonr Sandbox
            </Text>
          </Link>
          <Spacer />
          <Box>
            <Tag size="lg">v0.1.0</Tag>
          </Box>
        </Flex>
      }
    >
      <Box
        as="main"
        alignContent="center"
        marginLeft="20vw"
        marginRight="20vw"
        marginTop="15vh"
      >
        <Card title="Import Existing Sonr Account">
          <CardBody>
            <Form onSubmit={() => {}}>
              <Field type="text" name="address" label="Address" />
              {/* or: <PasswordField name="password" label="Password" /> */}
            </Form>
          </CardBody>
          <CardFooter>
            <ButtonGroup>
              <Link href="/">
                <Button label="Cancel" variant="ghost" />
              </Link>
            </ButtonGroup>
          </CardFooter>
        </Card>
      </Box>
    </AppShell>
  );
}
