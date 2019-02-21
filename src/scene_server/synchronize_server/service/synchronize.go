/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package service

import (
	"encoding/json"

	"github.com/emicklei/go-restful"

	"configcenter/src/common/metadata"
)

func (s *Service) Find(req *restful.Request, resp *restful.Response) {
	srvData := s.newSrvComm(req.Request.Header)
	input := &metadata.SynchronizeFetchInfoParameter{}
	if err := json.NewDecoder(req.Request.Body).Decode(input); err != nil {
		blog.Errorf("FindInstance , but decode body failed, err: %v,rid:%s", err, srvData.rid)
		resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: srvData.ccErr.Error(common.CCErrCommJSONUnmarshalFailed)})
		return
	}

	data, err := srvData.lgc.Find(srvData.ctx, input)
	if err != nil {
		blog.Errorf("FindInstance error. error: %s,input:%#v,rid:%s", err.Error(), input, srvData.rid)
		resp.WriteError(http.StatusInternalServerError, &meta.RespError{Msg: err})
		return
	}
	resp.WriteEntity(metadata.QueryConditionResult{
		BaseResp: meta.SuccessBaseResp,
		Data:     data,
	})
}
