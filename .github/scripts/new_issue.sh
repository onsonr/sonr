#!/bin/sh

ROOT_DIR=$(git rev-parse --show-toplevel)

select_scope() {
    cat "$ROOT_DIR/.github/scopes.json" | jq -r '.scopes[]' | fzf --prompt "Select scope:"
}

get_title() {
    gum input --placeholder "Issue Title..."
}

add_requirement() {
    requirement=$(gum input --placeholder "Add a requirement...")
    if [ -n "$requirement" ]; then
        REQUIREMENTS="$REQUIREMENTS
$requirement"
        return 0
    fi
    return 1
}

collect_requirements() {
    REQUIREMENTS=""
    req_count=0
    while true; do
        if add_requirement; then
            req_count=$((req_count + 1))
            if [ $req_count -ge 2 ] && ! gum confirm --default=false "Do you want to add another requirement?"; then
                break
            fi
        else
            if [ $req_count -ge 2 ]; then
                break
            else
                echo "Requirement cannot be empty. Please enter a valid requirement."
            fi
        fi
    done
}

get_docs() {
    docs=$(cat "$ROOT_DIR/.github/scopes.json" | jq -c '.docs')
     mods --role "determine-issue-docs" "$SCOPE" "$TITLE" "$docs"
}

get_goal() {
    mods --role "determine-issue-goal" "$SCOPE $TITLE"
}

format_requirements() {
    i=1
    echo "$REQUIREMENTS" | while IFS= read -r req; do
        if [ -n "$req" ]; then
            echo "$i. $req"
            i=$((i + 1))
        fi
    done
}

create_body() {
    goal=$(get_goal)
    docs=$(get_docs)
    
    echo "### Goal(s):"
    echo "$goal"
    echo
    echo "### Requirements:"
    format_requirements
    echo
    echo "### Resources:"
    echo "$docs"
}

preview_issue() {
    echo "# ($SCOPE) $TITLE"
    echo "$ISSUE_BODY"
}

create_github_issue() {
    draft_flag=""
    if gum confirm "Assign this issue to yourself?"; then
        draft_flag="-a @me"
    fi
    
    gh issue create \
        --repo onsonr/sonr \
        --title "[$SCOPE] $TITLE" \
        --body "$ISSUE_BODY" \
        $draft_flag
}

main() {
    SCOPE=$(select_scope)
    TITLE=$(get_title)
    collect_requirements
    ISSUE_BODY=$(create_body)
    
    preview_issue | gum format
    
    if gum confirm "Do you want to create a new GitHub issue with this information?"; then
        create_github_issue
    else
        exit 1
    fi
}

main

