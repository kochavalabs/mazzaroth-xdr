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
        BLOCK = 2 // Request a specific block from the node
    };

    union DownloadRequestPayload switch (DownloadRequestType Type)
    {
        case UNKNOWN:
            void;
        case HEIGHT:
            void;
        case BLOCK:
            unsigned hyper blockNumber;
    }

    // Download Responses returned from requests
    struct DownloadResponse {
        DownloadResponsePayload downloadResponsePayload;
    }

    union DownloadResponsePayload switch (DownloadRequestType Type)
    {
        case UNKNOWN:
            void;
        case HEIGHT:
            unsigned hyper height;
        case BLOCK:
            Block block;
    }
}