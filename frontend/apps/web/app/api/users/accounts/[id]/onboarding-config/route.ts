import { withNeosyncContext } from '@/api-only/neosync-context';
import { getSystemAppConfig } from '@/app/api/config/config';
import { RequestContext } from '@/shared';
import {
  GetAccountOnboardingConfigRequest,
  SetAccountOnboardingConfigRequest,
} from '@neosync/sdk';
import { NextRequest, NextResponse } from 'next/server';

export async function GET(
  req: NextRequest,
  { params }: RequestContext
): Promise<NextResponse> {
  const systemConfig = getSystemAppConfig();
  return withNeosyncContext(async (ctx) => {
    return ctx.client.users.getAccountOnboardingConfig(
      new GetAccountOnboardingConfigRequest({
        accountId: params.id,
      })
    );
  })(req);
}

export async function POST(req: NextRequest): Promise<NextResponse> {
  const systemConfig = getSystemAppConfig();
  return withNeosyncContext(async (ctx) => {
    return ctx.client.users.setAccountOnboardingConfig(
      SetAccountOnboardingConfigRequest.fromJson(await req.json())
    );
  })(req);
}

export async function PUT(req: NextRequest): Promise<NextResponse> {
  const systemConfig = getSystemAppConfig();
  return withNeosyncContext(async (ctx) => {
    return ctx.client.users.setAccountOnboardingConfig(
      SetAccountOnboardingConfigRequest.fromJson(await req.json())
    );
  })(req);
}
