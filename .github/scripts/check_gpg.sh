#!/usr/bin/env bash
set -euo pipefail

CHECKSUM_FILE="dist/checksums.txt"
SIGNATURE_FILE="dist/checksums.txt.sig"

echo "🔎 Validating checksum signature..."

# 1. Verifica se os arquivos existem
if [[ ! -f "$CHECKSUM_FILE" ]]; then
  echo "❌ Error: $CHECKSUM_FILE not found."
  exit 1
fi

if [[ ! -f "$SIGNATURE_FILE" ]]; then
  echo "❌ Error: $SIGNATURE_FILE not found."
  exit 1
fi

# 2. Verifica se o checksums.txt.sig é formato ASCII PGP
if ! grep -q "BEGIN PGP SIGNATURE" "$SIGNATURE_FILE"; then
  echo "❌ Error: $SIGNATURE_FILE is not a valid ASCII-armored PGP signature."
  exit 1
fi

# 3. Tenta verificar assinatura
if gpg --verify "$SIGNATURE_FILE" "$CHECKSUM_FILE"; then
  echo "✅ Signature is valid!"
else
  echo "❌ Invalid signature!"
  exit 1
fi