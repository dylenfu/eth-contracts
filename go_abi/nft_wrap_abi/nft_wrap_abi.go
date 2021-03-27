Fatal: Failed to build Solidity contract: solc: exit status 1
Error: Member "safeTransferFrom" not found or not visible after argument-dependent lookup in contract IERC721.
  --> ./contracts_nftlockproxy/PolyNFTWrap.sol:94:17:
   |
94 |         require(nftContract.safeTransferFrom(msg.sender, toAddress, tokenId, callData), "lock erc721 failed");
   |                 ^^^^^^^^^^^^^^^^^^^^^^^^^^^^


