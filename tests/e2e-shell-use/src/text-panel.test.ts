import { describe, it, beforeAll, afterAll } from 'vitest';
import { ShellUse } from '@microsoft/shell-use';

describe('Text Panel', () => {
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

  it('shows Lorem Ipsum text on the Text tab', async () => {
    await su.press('1');
    await su.waitIdle();
    await su.expectText('Lorem ipsum');
  });
});
