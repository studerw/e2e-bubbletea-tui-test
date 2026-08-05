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
    // Wait for the fetch to complete (3-7s simulated delay + network)
    await su.waitIdle({ timeout: 20000 });
    // The response should contain user data from JSONPlaceholder
    await su.expectText('Leanne Graham', { timeout: 20000 });
  });
});
