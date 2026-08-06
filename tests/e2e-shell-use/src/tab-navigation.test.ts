import { describe, it, beforeAll, afterAll } from 'vitest';
import { ShellUse } from '@microsoft/shell-use';

describe('Tab Navigation', () => {
  let su: ShellUse;

  beforeAll(async () => {
    su = new ShellUse();
    await su.run('../../bin/tui-e2e-demo');
    await su.waitIdle();
  });

  afterAll(async () => {
    try {
      await su.close();
    } catch {
      // Session may already be closed
    }
  });

  it('switches to Web Call tab with key 2', async () => {
    await su.press('2');
    await su.waitIdle();
    // Web Call panel should show endpoint options
    await su.expectText('User #1');
  });

  it('switches to Todo List tab with key 3', async () => {
    await su.press('3');
    await su.waitIdle();
    // Assert on a pre-loaded todo item — unique to the Todo List panel content
    await su.expectText('Write E2E tests for TUI app');
  });

  it('switches to Timer tab with key 4', async () => {
    await su.press('4');
    await su.waitIdle();
    // Assert on the initial timer display — unique to the Timer panel
    await su.expectText('00:00');
  });

  it('switches back to Text tab with key 1', async () => {
    await su.press('1');
    await su.waitIdle();
    await su.expectText('Lorem ipsum');
  });

  it('cycles tabs with Tab key', async () => {
    // Starting at tab 1 (Text), pressing Tab should go to tab 2 (Web Call)
    await su.press('Tab');
    await su.waitIdle();
    await su.expectText('User #1');
  });
});
