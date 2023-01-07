import { Box, Center, Flex, Spacer, Tag, Text } from "@chakra-ui/react";
import {
  FiKey,
  FiDroplet,
  FiUserPlus,
  FiAtSign,
  FiCloud,
  FiLock,
} from "react-icons/fi";

import {
  AppShell,
  Button,
  ButtonGroup,
  Card,
  CardBody,
  Field,
  FormLayout,
  FormStep,
  FormStepper,
  Loader,
  NextButton,
  PrevButton,
  Property,
  PropertyList,
  StepForm,
  StepperCompleted,
  useSnackbar,
} from "@saas-ui/react";
import { Web3Address } from "@saas-ui/web3";
import Link from "next/link";
import { useState } from "react";
import * as Yup from "yup";

export default function SignUp() {
  const snackbar = useSnackbar();
  const [loading, setLoading] = useState(false);
  const [label, updateLabel] = useState("");
  const [address, setAddress] = useState("");
  const [vaultCid, setVaultCid] = useState("");
  const [cmpConfig, setCmpConfig] = useState("");
  const [session, setSession] = useState("");
  const [credential, setCredential] = useState<PublicKeyCredential | null>(
    null
  );
  const schemas = {
    credential: Yup.object().shape({
      deviceLabel: Yup.string().required().label("Device Label"),
    }),
  };

  const yupResolver = (schema: any) => async (values: any) => {
    return {
      values,
    };
  };

  const onSubmit = (params: any) => {
    console.log(params);
    return new Promise((resolve) => {
      setTimeout(resolve, 1000);
    });
  };

  const createCredential = async (nextStep: () => void) => {
    const response = await fetch("/api/vault/challenge", {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });

    const data = await response.json();
    const challenge = data.challenge as string;
    const sessionId = data.session_id as string;
    const credential = await navigator.credentials.create({
      publicKey: {
        rp: {
          name: "Sonr",
        },
        user: {
          id: Uint8Array.from(sessionId, (c) => c.charCodeAt(0)),
          name: label,
          displayName: label,
        },
        challenge: Uint8Array.from(challenge, (c) => c.charCodeAt(0)),
        pubKeyCredParams: [
          {
            type: "public-key",
            alg: -7,
          },
        ],
        timeout: 60000,
        attestation: "direct",
      },
    });
    setCredential(credential as PublicKeyCredential);
    setSession(sessionId);
    snackbar.info("PassKey has been generated by your device.");
    nextStep();
  };

  const registerAccount = async function callVaultRegister(
    nextStep: () => void
  ): Promise<void> {
    setLoading(true);
    const response = await fetch("/api/vault/register", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        session_id: session,
        credential_response: credential?.response.clientDataJSON,
      }),
    });
    console.log(response);

    // Get response as object
    const data = await response.json();
    if (data.success) {
      snackbar.info("Account has been registered with Sonr Vault.");
      setLoading(false);
      nextStep();
    } else {
      snackbar.error("Account registration failed.");
      setLoading(false);
      return;
    }
  };

  const accountKeygen = async function callVaultKeygen(
    nextStep: () => void
  ): Promise<string> {
    setLoading(true);
    const response = await fetch("/api/vault/keygen", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
    });
    // Get response as object
    const data = await response.json();
    setAddress(data.address);
    setVaultCid(data.vault_cid);
    setCmpConfig(data.share_config.cmp_config);
    setLoading(false);
    snackbar.info("Account has been registered with Sonr Vault.");
    nextStep();
    return data.address;
  };

  return (
    <AppShell
      maxWidth="100vw"
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
            <Tag size="lg">{process.env.NEXT_PUBLIC_VERCEL_GIT_COMMIT_REF}</Tag>
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
            maxWidth="600px"
            variant="solid"
            title="Sonr.ID"
            subtitle="Use the Vault MPC Protocol with Webauthn."
            padding={4}
            action={
              <ButtonGroup>
                <Link href="/">
                  <Button variant="unstyled" colorScheme="red">
                    Cancel
                  </Button>
                </Link>
              </ButtonGroup>
            }
          >
            <CardBody margin={4}>
              <StepForm
                defaultValues={{
                  deviceLabel: "",
                  credentialId: "",
                  credentialType: "",
                  credentialPublicKey: "",
                  credentialResponse: "",
                }}
                onSubmit={onSubmit}
              >
                {({
                  isFirstStep,
                  isLastStep,
                  isCompleted,
                  nextStep,
                  prevStep,
                }) => (
                  <FormLayout>
                    <FormStepper orientation="vertical">
                      <FormStep
                        name="project"
                        title="Generate PassKey"
                        resolver={yupResolver(schemas.credential)}
                      >
                        <FormLayout>
                          <Field
                            isRequired
                            name="deviceLabel"
                            label="Label"
                            onInput={(event) => {
                              // Check if the input is a string
                              const str = event.target as HTMLInputElement;

                              // Get the value from the input
                              const value = str.value;
                              updateLabel(str.value);
                            }}
                          />
                          <Button
                            leftIcon={<FiKey />}
                            label="New PassKey"
                            onClick={() =>
                              label
                                ? createCredential(nextStep)
                                : snackbar.error("Please enter a device label")
                            }
                            variant="outline"
                          />
                        </FormLayout>
                      </FormStep>

                      <FormStep name="register" title="Register Account">
                        {loading ? (
                          <Flex>
                            <Loader>Running MPC Protocol...</Loader>
                          </Flex>
                        ) : (
                          <FormLayout>
                            <PropertyList>
                              <Property
                                label="Label"
                                value={label ? label : "No credential"}
                              />
                              <Property
                                label="Credential ID"
                                value={
                                  <Web3Address
                                    address={
                                      credential
                                        ? credential.id
                                        : "No credential"
                                    }
                                    startLength={credential ? 12 : 15}
                                    endLength={credential ? 4 : 0}
                                  />
                                }
                              />
                              <Property
                                label="Type"
                                value={
                                  credential ? credential.type : "No credential"
                                }
                              />
                              <Property label="Source" value="WebAuthn" />
                            </PropertyList>
                            <ButtonGroup>
                              <Button
                                leftIcon={<FiUserPlus />}
                                label="Register Account"
                                onClick={() => registerAccount(nextStep)}
                              />
                            </ButtonGroup>
                          </FormLayout>
                        )}
                      </FormStep>

                      <FormStep name="faucet" title="Get Tokens from Faucet">
                        <FormLayout>
                          <Text>
                            Please confirm that your information is correct.
                          </Text>
                          <PropertyList>
                            <Property label="Address" />
                            <Button variant="outline" leftIcon={<FiAtSign />}>
                              <Web3Address
                                address={address ? address : "N/A"}
                                startLength={32}
                                endLength={4}
                              />
                            </Button>
                            <Property label="Vault CID" />
                            <Button variant="outline" leftIcon={<FiCloud />}>
                              <Web3Address
                                address={vaultCid ? vaultCid : "N/A"}
                                startLength={32}
                                endLength={4}
                              />
                            </Button>

                            <Property label="Wallet Share" />
                            <Box>
                              <Button variant="outline" leftIcon={<FiLock />}>
                                <Web3Address
                                  address={cmpConfig ? cmpConfig : "N/A"}
                                  startLength={32}
                                  endLength={4}
                                />
                              </Button>
                            </Box>
                          </PropertyList>
                          <ButtonGroup>
                            <Button
                              leftIcon={<FiDroplet />}
                              label="Get Airdrop"
                              onClick={() =>
                                snackbar.error(
                                  "Faucet not online at this time. Please try again later."
                                )
                              }
                              variant="primary"
                            />
                          </ButtonGroup>
                        </FormLayout>
                      </FormStep>
                      <FormStep
                        name="confirm"
                        title="Broadcast Document Transaction"
                      >
                        <FormLayout>
                          <Text>
                            Please confirm that your information is correct.
                          </Text>
                          <ButtonGroup>
                            <NextButton />
                            <PrevButton variant="ghost" />
                          </ButtonGroup>
                        </FormLayout>
                      </FormStep>

                      <StepperCompleted>
                        <Loader>
                          We are setting up your project, just a moment...
                        </Loader>
                      </StepperCompleted>
                    </FormStepper>
                  </FormLayout>
                )}
              </StepForm>
            </CardBody>
          </Card>
        </Box>
      </Center>
    </AppShell>
  );
}
