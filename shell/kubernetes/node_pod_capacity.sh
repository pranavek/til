#!/bin/bash
# Script to show pod usage and capacity across all nodes

export KUBECONFIG=~/.atmos2/config-atmos-scratch-avondale

echo "============================================================================================================"
echo "NODE POD CAPACITY REPORT - $(date)"
echo "============================================================================================================"
printf "%-45s %-12s %-12s %-8s %-15s\n" "NODE NAME" "PODS (USED)" "CAPACITY" "USAGE %" "AGE"
echo "------------------------------------------------------------------------------------------------------------"

# Get all nodes
nodes=$(kubectl get nodes -o jsonpath='{.items[*].metadata.name}')

total_used=0
total_capacity=0

for node in $nodes; do
    # Get pod capacity
    capacity=$(kubectl get node "$node" -o jsonpath='{.status.allocatable.pods}')

    # Count actual pods on the node
    used=$(kubectl get pods --all-namespaces --field-selector spec.nodeName="$node" --no-headers 2>/dev/null | wc -l)

    # Calculate percentage
    if [ "$capacity" -gt 0 ]; then
        percentage=$((used * 100 / capacity))
    else
        percentage=0
    fi

    # Get node age
    age=$(kubectl get node "$node" -o jsonpath='{.metadata.creationTimestamp}' | xargs -I {} date -d {} +"%Y-%m-%d %H:%M" 2>/dev/null || echo "N/A")

    # Color coding based on usage
    if [ "$percentage" -ge 95 ]; then
        status="🔴 CRITICAL"
    elif [ "$percentage" -ge 80 ]; then
        status="🟡 HIGH"
    else
        status="🟢 OK"
    fi

    # Short node name (remove domain suffix)
    short_node=$(echo "$node" | cut -d. -f1)

    printf "%-45s %-12s %-12s %-8s %-15s %s\n" \
        "$short_node" \
        "$used" \
        "$capacity" \
        "${percentage}%" \
        "$age" \
        "$status"

    total_used=$((total_used + used))
    total_capacity=$((total_capacity + capacity))
done

echo "------------------------------------------------------------------------------------------------------------"

if [ "$total_capacity" -gt 0 ]; then
    total_percentage=$((total_used * 100 / total_capacity))
else
    total_percentage=0
fi

printf "%-45s %-12s %-12s %-8s\n" \
    "CLUSTER TOTAL" \
    "$total_used" \
    "$total_capacity" \
    "${total_percentage}%"

echo "============================================================================================================"
