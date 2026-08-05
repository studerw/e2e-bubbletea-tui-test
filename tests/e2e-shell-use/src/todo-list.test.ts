import { describe, it, beforeAll, afterAll } from 'vitest';
import { ShellUse } from '@microsoft/shell-use';

describe('Todo List Panel', () => {
  let su: ShellUse;

  beforeAll(async () => {
    su = new ShellUse();
    await su.run('../../bin/tui-e2e-demo');
    await su.waitIdle();
    // Navigate to Todo List tab
    await su.press('3');
    await su.waitIdle();
  });

  afterAll(async () => {
    try {
      await su.close();
    } catch {
      // Session may already be closed
    }
  });

  it('adds a new todo item', async () => {
    // Press 'a' to enter input mode
    await su.press('a');
    await su.waitIdle();
    // Type the new item
    await su.type('New E2E Item');
    await su.waitIdle();
    // Press Enter to submit
    await su.press('Enter');
    await su.waitIdle();
    // Verify the new item appears
    await su.expectText('New E2E Item');
  });
});
