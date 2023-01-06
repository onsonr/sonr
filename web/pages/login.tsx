import { Box, Divider, Text } from "@chakra-ui/react";
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
        <>
          <Box
            as="header"
            alignItems="start"
            borderBottomWidth="1px"
            py="2"
            px="4"
          >
            <Link href="/">
              <Button variant="outline">Return Home</Button>
            </Link>
          </Box>
        </>
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
            <ButtonGroup paddingTop={8}>
              <Button variant="outline">Cancel</Button>
              <Button variant="primary">Import</Button>
            </ButtonGroup>
          </CardFooter>
        </Card>
      </Box>
    </AppShell>
  );
}
