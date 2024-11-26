import { OpenFeatureProvider } from "@openfeature/react-sdk";
import { Flowbite } from "flowbite-react";
import * as log from "loglevel";
import type { FC } from "react";
import { useEffect } from "react";
import { initGoFeatureFlag } from "../../helpers/goff.ts";
import { flowbiteTheme } from "../../theme.ts";
import { SidebarProvider } from "../navigation/SidebarContext.tsx";
import { NavigationLayoutContent } from "../navigation/navigationLayoutContent.tsx";
import type { LayoutProps } from "./layoutProps.ts";

export const RootLayout: FC<LayoutProps> = ({ children }) => {
  useEffect(() => {
    const initializeFeatureFlag = async () => {
      await initGoFeatureFlag();
    };
    initializeFeatureFlag().catch((error) => {
      log.error("error when initializing OpenFeature", error);
    });
  }, []);
  return (
    <>
      <OpenFeatureProvider>
        <Flowbite theme={{ theme: flowbiteTheme }}>
          <SidebarProvider>
            <NavigationLayoutContent>{children}</NavigationLayoutContent>
          </SidebarProvider>
        </Flowbite>
      </OpenFeatureProvider>
    </>
  );
};