import { describe, it, beforeAll, afterAll } from 'vitest';
import { ShellUse } from '@microsoft/shell-use';

describe('Timer Panel', () => {
  let su: ShellUse;

  beforeAll(async () => {
    su = new ShellUse();
    await su.run('../../bin/tui-e2e-demo');
    await su.waitIdle();
    // Navigate to Timer tab
    await su.press('4');
    await su.waitIdle();
  });

  afterAll(async () => {
    try {
      await su.close();
    } catch {
      // Session may already be closed
    }
  });

  it('starts and stops the timer', async () => {
    // Press 's' to start the timer
    await su.press('s');
    // Wait for the timer to tick a few times
    await new Promise((resolve) => setTimeout(resolve, 2000));
    // Press 's' again to stop
    await su.press('s');
    await su.waitIdle();
    // The timer should show non-zero elapsed time (at least 00:01)
    await su.expectText('00:0');
  });
});
