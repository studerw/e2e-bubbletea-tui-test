import { defineConfig } from 'vitest/config';

export default defineConfig({
  test: {
    testTimeout: 60000, // 60s per test for Shell-Use session startup + operations
    hookTimeout: 30000, // 30s for beforeAll/afterAll
    sequence: {
      concurrent: false, // Run tests sequentially
    },
    include: ['src/**/*.test.ts'],
  },
});
