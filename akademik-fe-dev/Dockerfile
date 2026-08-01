# Stage 1: Install dependencies
FROM node:lts-alpine3.23 AS deps
RUN apk add --no-cache libc6-compat
WORKDIR /app

COPY package.json package-lock.json* ./
RUN npm ci

# Stage 2: Rebuild the source code
FROM node:lts-alpine3.23 AS builder
WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
COPY . .

ARG NEXT_PUBLIC_API_BASE_URL
ARG NEXT_PUBLIC_API_DATA_REFERENSI_URL
ARG NEXT_PUBLIC_API_SSO_URL
ARG NEXT_PUBLIC_API_SDM_URL
ARG NEXT_PUBLIC_UI_SSO_URL

ENV NEXT_PUBLIC_API_BASE_URL=$NEXT_PUBLIC_API_BASE_URL
ENV NEXT_PUBLIC_API_DATA_REFERENSI_URL=$NEXT_PUBLIC_API_DATA_REFERENSI_URL
ENV NEXT_PUBLIC_API_SSO_URL=$NEXT_PUBLIC_API_SSO_URL
ENV NEXT_PUBLIC_API_SDM_URL=$NEXT_PUBLIC_API_SDM_URL
ENV NEXT_PUBLIC_UI_SSO_URL=$NEXT_PUBLIC_UI_SSO_URL

RUN npm run build

# Stage 3: Production image
FROM node:lts-alpine3.23 AS runner
WORKDIR /app

ENV NODE_ENV production

RUN addgroup --system --gid 1001 nodejs
RUN adduser --system --uid 1001 nextjs

COPY --from=builder /app/public ./public
COPY --from=builder --chown=nextjs:nodejs /app/.next/standalone ./
COPY --from=builder --chown=nextjs:nodejs /app/.next/static ./.next/static

USER nextjs

EXPOSE 3000

ENV PORT 3000

CMD ["node", "server.js"]
