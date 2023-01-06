import Head from "next/head";
import Image from "next/image";
import { Web3Address } from "@saas-ui/web3";
import styles from "../styles/Home.module.css";
import { useState } from "react";
import { Box, Center } from "@chakra-ui/react";
import { AppShell } from "@saas-ui/app-shell";
import { SignUp } from "./signup";
import {
  Button,
  Card,
  CardContainer,
  CardHeader,
  CardTitle,
  EmptyState,
  MenuButton,
  MenuItem,
  CardMedia,
  CardBody,
  CardFooter,
} from "@saas-ui/react";
import { ResponsiveMenu, ResponsiveMenuList } from "@saas-ui/pro";

export default function Home() {
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
          <>
            <Box as="header" align="end" borderBottomWidth="1px" py="2" px="4">
              <ResponsiveMenu>
                <MenuButton as={Button}>No Accounts</MenuButton>
                <ResponsiveMenuList>
                  <MenuItem>Edit</MenuItem>
                  <MenuItem>Delete</MenuItem>
                </ResponsiveMenuList>
              </ResponsiveMenu>
            </Box>
          </>
        }
      >
        <Center>
          <>
            <EmptyState
              colorScheme="primary"
              // icon={FiUsers}
              title="No Accounts Yet"
              description="You haven't imported any customers yet."
              actions={
                <>
                  <Button
                    label="Import account"
                    colorScheme="primary"
                    href="/login"
                  />
                  <Button label="Create account" href="/signup" />
                </>
              }
            />
          </>
        </Center>
      </AppShell>
    </>
  );
}
