#!/bin/bash 
echo "Creating Proposal For Upgrade"
echo "Enter Title:"
read -r title
echo "Enter Height:"
read -r height
echo "Enter From User Address:"
read -r from
echo "Submit Proposal To Upgrade";
# Submit Proposal
sonrd tx gov submit-proposal software-upgrade --upgrade-height $height --from $from --yes --title $title --description test upgrade
echo "Enter the Proposal ID:"
read -r proposal_id
# Deposit For Proposal
echo "Depositing Funds For Proposal ID: $proposal_id"
sonrd tx gov deposit $proposal_id 10000000stake --from $from --yes
# Vote for the Proposal
echo "Voting Yes For Proposal ID: $proposal_id"
sonrd tx gov vote $proposal_id yes --from $from --yes