---
description: >-
  Launching a protocol and its token holds significant importance in
  authorities' assessment of the protocol's decentralization.
---

# Decentralized Protocol Launch

Launching a protocol and its token holds significant importance in authorities' assessment of the protocol's decentralization. A key question often asked is: Who was responsible for minting and distributing the token?Many recently launched networks consider it crucial to avoid a scenario where the founders or core team take sole responsibility for minting the network's tokens during its launch. In the case of Cosmos SDK-based blockchains, **decentralization is inherent by default**.

Although the founding team typically creates the genesis file, which contains information about token minting and distribution, they do not have the authority to decide whether the network should be launched. Validators play a vital role by initiating the launch through signing the genesis file and launching the network. To effectively create and distribute the majority of tokens, a joint decision must be made by two-thirds of the voting power in the genesis file.

During this phase, token allocation occurs based on the information specified in the genesis file. The file outlines the initial distribution of tokens to various stakeholders, including the founding team, investors, airdrop recipients, the community, and the network's treasury.

### **Disabling IBC & Token Transfers**

During the initial launch, there is the option to **temporarily disable Inter-Blockchain Communication (IBC) token transfers** and **Internal Token transfers**. This serves as a precautionary measure to establish a controlled and secure environment for the mainnet's early stages, often referred to as a "Soft Launch."

Another benefit is that disabling IBC and token transfers prevent the **immediate creation of secondary markets**. This approach ensures that the community establishes secondary markets only after enabling IBC and token transfers, aligning them with the project's objectives and regulatory considerations.

Once the soft launch proves successful, all features are thoroughly reviewed, and a more decentralized distribution of voting power is achieved, the community can then enable IBC and token transfers through a governance parameter change proposal. Involving the community in decision-making reduces potential liabilities for both the founders and validators involved in the network’s launch. This, in turn, increases the decentralization of the decision-making process regarding the creation of secondary markets.
