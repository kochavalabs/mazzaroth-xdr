%#include "block.x"

namespace mazzaroth
{
    // Download Requests are made between Mazzaroth nodes to sync data
    struct DownloadRequest {
        DownloadRequestPayload downloadRequestPayload;
    }

    // The different request messages that can be sent
    enum DownloadRequestType {
        UNKNOWN = 0,
        HEIGHT = 1, // Request current block height of node
        BLOCK = 2, // Request a specific block from the node
        BATCHES = 3 // Request the missing batches after last sequence number
    };

    union DownloadRequestPayload switch (DownloadRequestType Type)
    {
        case UNKNOWN:
            void;
        case HEIGHT:
            void;
        case BLOCK:
            unsigned hyper blockNumber;
        case BATCHES:
            BatchesRequest batchesRequest;
    }

    struct BatchesRequest {
        unsigned hyper seqNum;
        // Node details to target peer with batches
        string id;
        string ip;
        unsigned hyper port;
    }

    // Download Responses returned from requests
    struct DownloadResponse {
        DownloadStatus downloadStatus;
        DownloadResponsePayload downloadResponsePayload;
    }

    // Status of a download request.
    enum DownloadStatus
    {
        // The status is either not known or not set.
        UNKNOWN = 0,

        // Download request was successful.
        SUCCESS = 1,

        // Download request failed.
        FAILURE = 2,
    };

    union DownloadResponsePayload switch (DownloadRequestType Type)
    {
        case UNKNOWN:
            void;
        case HEIGHT:
            DownloadHeight downloadHeight;
        case BLOCK:
            Block block; // Return the requested block
        case BATCHES:
            void; // Missing batches are sent in new messages, not a response
    }

    // Download Height info including ledger block height and last processed sequence number
    struct DownloadHeight {
        unsigned hyper height; // Return ledger height
        unsigned hyper seqNum; // Return latest consensus sequence number
    }
}