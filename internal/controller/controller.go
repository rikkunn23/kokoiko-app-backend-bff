package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/rikkunn23/kokoiko-app-backend-bff/internal/logger"
	"github.com/rikkunn23/kokoiko-app-backend-bff/pkg/werrors"
)

// レスポンスの内容次第で返答するHTTPステータスコードを決定とかする
func ResponseHandler(ctx context.Context, w http.ResponseWriter, res any, err error) {
  w.Header().Set("Content-Type", "application/json")

	// エラーがnilでない、エラーがある場合はエラーハンドリングを行う
  if err != nil {
    handleError(ctx, w, err)
    return
  }
	// レスポンスがnil、成功だけど返すものがない場合は204 No Contentを返却
  // bodyが無い場合は204を返却
  if res == nil {
    w.WriteHeader(http.StatusNoContent)
    return
  }
  if err := json.NewEncoder(w).Encode(res); err != nil {
    logger.ErrorfWithCtx(ctx, "json encode error: %+v\n", err)
    return
  }
}

func handleError(ctx context.Context, w http.ResponseWriter, err error) {
  switch errType := err.(type) {
  case werrors.InternalError:
    logger.WerrorWithCtx(ctx, errType)
    w.WriteHeader(http.StatusInternalServerError)
    _ = json.NewEncoder(w).Encode(werrors.InternalErrorResponse{
      BaseErrorResponse: werrors.BaseErrorResponse{
        Code: errType.Code,
        Msg:  errType.Error(),
      },
    })
  case werrors.InvalidArgumentError:
    logger.WwarnWithCtx(ctx, errType)
    w.WriteHeader(http.StatusBadRequest)
    _ = json.NewEncoder(w).Encode(werrors.InvalidArgumentErrorResponse{
      BaseErrorResponse: werrors.BaseErrorResponse{
        Code: errType.Code,
        Msg:  errType.Error(),
      },
    })
  case werrors.UnauthorizedError:
    // tokenの有効期限切れはwarnにする必要がないのでinfoに落としている
    if errType.Error() == "expired_token" {
      logger.WinfoWithCtx(ctx, errType)
    } else {
      logger.WwarnWithCtx(ctx, errType)
    }
    w.WriteHeader(http.StatusUnauthorized)
    _ = json.NewEncoder(w).Encode(werrors.UnauthorizedErrorResponse{
      BaseErrorResponse: werrors.BaseErrorResponse{
        Code: errType.Code,
        Msg:  errType.Error(),
      },
      Detail: errType.Detail,
    })
  case werrors.ForbiddenError:
    logger.WwarnWithCtx(ctx, errType)
    w.WriteHeader(http.StatusForbidden)
    _ = json.NewEncoder(w).Encode(werrors.ForbiddenErrorResponse{
      BaseErrorResponse: werrors.BaseErrorResponse{
        Code: errType.Code,
        Msg:  errType.Error(),
      },
    })
  case werrors.NotFoundError:
    logger.WwarnWithCtx(ctx, errType)
    w.WriteHeader(http.StatusNotFound)
    _ = json.NewEncoder(w).Encode(werrors.NotFoundErrorResponse{
      BaseErrorResponse: werrors.BaseErrorResponse{
        Code: errType.Code,
        Msg:  errType.Error(),
      },
    })
  case werrors.ConflictError:
    logger.WwarnWithCtx(ctx, errType)
    w.WriteHeader(http.StatusConflict)
    _ = json.NewEncoder(w).Encode(werrors.ConflictErrorResponse{
      BaseErrorResponse: werrors.BaseErrorResponse{
        Code: errType.Code,
        Msg:  errType.Error(),
      },
    })
  default:
    logger.ErrorWithCtx(ctx, err)
    w.WriteHeader(http.StatusInternalServerError)
    _ = json.NewEncoder(w).Encode(werrors.InternalErrorResponse{
      BaseErrorResponse: werrors.BaseErrorResponse{
        Code: http.StatusInternalServerError,
        Msg:  errType.Error(),
      },
    })
  }
}
