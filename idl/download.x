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
        PENDING = 3 // Request the pending transactions in ledger
    };

    union DownloadRequestPayload switch (DownloadRequestType Type)
    {
        case UNKNOWN:
            void;
        case HEIGHT:
            void;
        case BLOCK:
            unsigned hyper blockNumber;
        case PENDING:
            void;
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
            unsigned hyper height;
        case BLOCK:
            Block block;
        case PENDING:
            DownloadPendingPayload pendingPayload;
    }

    // Download Pending Response includes the pending transactions
    // and current ledger height
    struct DownloadPendingPayload {
        unsigned hyper height;
        Transaction transactions<>;
    }
}