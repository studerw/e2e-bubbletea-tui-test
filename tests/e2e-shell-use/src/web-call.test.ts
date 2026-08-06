import { describe, it, beforeAll, afterAll } from 'vitest';
import { ShellUse } from '@microsoft/shell-use';

describe('Web Call Panel', () => {
  let su: ShellUse;

  beforeAll(async () => {
    su = new ShellUse();
    await su.run('../../bin/tui-e2e-demo');
    await su.waitIdle();
    // Navigate to Web Call tab
    await su.press('2');
    await su.waitIdle();
  });

  afterAll(async () => {
    try {
      await su.close();
    } catch {
      // Session may already be closed
    }
  });

  it('fetches data from the selected endpoint', async () => {
    // Select the first endpoint (User #1) and press Enter to fetch
    await su.press('Enter');
    // The panel has a 3-7s simulated delay + spinner animation.
    // waitIdle will block until the spinner stops (i.e. fetch completes),
    // then expectText checks the response is visible.
    // Assert on JSON content that appears near the top of the response —
    // 'address' sorts first alphabetically so it's always visible and proves
    // real data was returned from the API.
    await su.expectText('"address": {', { timeout: 20000 });
  });
});
