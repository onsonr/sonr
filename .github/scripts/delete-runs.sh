#!/bin/zsh

org=sonr-io
repo=core

# Get workflow IDs with status "disabled_manually"
workflow_ids=($(gh api repos/$org/$repo/actions/workflows | jq '.workflows[] | .id'))

for workflow_id in "${workflow_ids[@]}"
do
  echo "Listing runs for the workflow ID $workflow_id"
  run_ids=( $(gh api repos/$org/$repo/actions/workflows/$workflow_id/runs | jq '.workflow_runs[].id') )
  for run_id in "${run_ids[@]}"
  do
    echo "Deleting Run ID $run_id"
    gh api repos/$org/$repo/actions/runs/$run_id -X DELETE >/dev/null
  done
done
