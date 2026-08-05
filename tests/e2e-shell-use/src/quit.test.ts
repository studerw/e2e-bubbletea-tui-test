import { describe, it, beforeAll } from 'vitest';
import { ShellUse } from '@microsoft/shell-use';

describe('Quit Behavior', () => {
  let su: ShellUse;

  beforeAll(async () => {
    su = new ShellUse();
    await su.run('../../bin/tui-e2e-demo');
    await su.waitIdle();
  });

  it('exits cleanly when q is pressed', async () => {
    await su.press('q');
    await su.waitExit();
  });
});
