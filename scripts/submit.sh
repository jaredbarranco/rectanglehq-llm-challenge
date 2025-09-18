# This uses the x-request-id from the successful start of stage 2.
REQUEST_ID=$(cat ./final_request.txt)
EMAIL="barrancojared+llm@gmail.com"
STRING_TO_HASH="${EMAIL}:${REQUEST_ID}"

# Step 5: Compute the hash.
# HASH=$(echo -n "$STRING_TO_HASH" | openssl dgst -blake2b512 | sed 's/^.* //')
HASH=$(echo -n "$STRING_TO_HASH" | b2sum -l 512 | awk '{print $1}')
echo $HASH
# Step 6: Print the final curl command.
echo "JSON payload created in final_submission.json."
echo "Run the following command to submit your answer:"
echo ""
echo "curl -i -L -X POST \\"
echo "  -H \"X-Application-Key: FDE-CHALLENGE-01\" \\"
echo "  -H \"Content-Type: application/json\" \\"
echo "  -H \"X-Idempotency-Key: ${HASH}\" \\"
echo "  -d @final_post.json \\"
echo "  challenge.rectanglehq.com"

