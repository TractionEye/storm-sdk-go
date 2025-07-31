package error

import "github.com/pingcap/errors"

var ErrAmmActionPhaseBounce = errors.New("amm action phase bounce")

type ContractError = int

const (
	ErrNotAnAdmin             ContractError = 401
	ErrMinGas                 ContractError = 404
	ErrWrongPositionAddress   ContractError = 405
	ErrFundingTime            ContractError = 406
	ErrNotAVault              ContractError = 407
	ErrLessMarginRatio        ContractError = 409
	ErrWrongDirection         ContractError = 410
	ErrWrongSize              ContractError = 411
	ErrPositionBadDebt        ContractError = 412
	ErrWrongPositionTimestamp ContractError = 413
	ErrWrongAmount            ContractError = 414
	ErrHighPriceImpact        ContractError = 415
	ErrOverSpreadLimit        ContractError = 416
	ErrOverMaxOpenNotional    ContractError = 417
	ErrNotExecutableByStop    ContractError = 418
	ErrNotExecutable          ContractError = 419
	ErrWrongBaseAssetReserve  ContractError = 420
	ErrWrongQuoteAssetReserve ContractError = 421
	ErrWrongID                ContractError = 422
	ErrQuorumFailed           ContractError = 423
	ErrInvalidTimestamp       ContractError = 424
	ErrInvalidDiff            ContractError = 425
	ErrWrongLeverage          ContractError = 426
	ErrSlippageTolerance      ContractError = 427
	ErrNegativeMarginToTrader ContractError = 428
	ErrOrderExpired           ContractError = 429
	ErrWrongLiquidity         ContractError = 430
	ErrInvalidBaseAssetAmount ContractError = 431
	ErrPaused                 ContractError = 432
	ErrCloseOnly              ContractError = 433
	ErrMarketIsActive         ContractError = 434
	ErrMarketIsNotActive      ContractError = 435
	ErrWrongReferralItem      ContractError = 436
	ErrWrongExecutorItem      ContractError = 437

	ErrStartAlreadySet        ContractError = 438
	ErrEndAlreadySet          ContractError = 439
	ErrNotStarted             ContractError = 440
	ErrEnded                  ContractError = 441
	ErrWrongStartDate         ContractError = 442
	ErrWrongEndDate           ContractError = 443
	ErrWrongExecutorAddress   ContractError = 444
	ErrDirectIncreaseDisabled ContractError = 445
	ErrDirectCloseDisabled    ContractError = 446

	ErrUnknownOraclePayload     ContractError = 447
	ErrWrongSingleOraclePayload ContractError = 448
	ErrWrongDoubleOraclePayload ContractError = 449

	ErrWrongPause                ContractError = 450
	ErrMarketPaused              ContractError = 451
	ErrNotEnoughPnl              ContractError = 452
	ErrHighCumulativePriceImpact ContractError = 453
	ErrPauseAfterUnpause         ContractError = 454
	ErrTooCloseToLiquidation     ContractError = 455
	ErrUnpauseNotDefined         ContractError = 456
	ErrInvalidPauseAt            ContractError = 457
	ErrInvalidUnpauseAt          ContractError = 458

	ErrWrongOp ContractError = 0xFFFF

	ErrFeeNotProvided = 901
)

func NewContractErr(e ContractError) error {
	return errors.Wrapf(ErrAmmActionPhaseBounce, "exit code: %d", e)
}
