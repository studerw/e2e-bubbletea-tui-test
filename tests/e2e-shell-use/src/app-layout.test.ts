import { describe, it, beforeAll, afterAll } from 'vitest';
import { ShellUse } from '@microsoft/shell-use';

describe('App Layout', () => {
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

  it('displays the header', async () => {
    await su.expectText('E2E MVP using BubbleTea v2');
  });

  it('displays the footer', async () => {
    await su.expectText('Clarity Innovations');
  });

  it('displays tab labels', async () => {
    await su.expectText('Text');
    await su.expectText('Web Call');
    await su.expectText('Todo List');
    await su.expectText('Timer');
  });

  it('displays initial panel content', async () => {
    // The Text panel (tab 1) should be shown by default
    await su.expectText('Lorem ipsum');
  });
});
