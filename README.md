**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/feitianyul/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/feitianyul/free-proxy-list.svg">
  
  <br>
  <br>
  <a href="https://getfreeproxy.com/lists/" title="可用代理列表">可用代理列表</a> | <a href="https://getfreeproxy.com/tools/proxy-checker" title="在线代理检测">免费代理检测</a> | <a href="https://getfreeproxy.com/tools/proxy-protocol-parser" title="代理协议解析">通用代理协议解析</a> | <a href="https://developer.getfreeproxy.com/" title="代理 API">免费代理 API</a>
  <br>
</p>

# 🌎 GetFreeProxy (GFP)：免费代理列表

**GetFreeProxy (GFP)** 是一个开源项目，自动从互联网聚合并校验免费代理，旨在为开发者、研究人员及需要代理服务的用户提供新鲜、可靠、可用的公共代理列表。

列表按小时更新，确保您始终能获取到最新的可用代理。

---

## 📖 项目说明

本项目为开源免费代理聚合与校验工具，从互联网公开源拉取代理并**仅保留 HTTP、HTTPS** 两种类型，经校验后生成列表，供开发者、研究人员等使用。

### 本仓库特点

- **仅保留两种代理**：HTTP、HTTPS，不收录 SOCKS、VMess、Trojan、VLESS、SS/SSR、Hysteria 等其它协议。
- **校验规则**：五域名中**任意 3 个**在 2 秒内成功（HTTP 200）即视为该协议通过。对每条代理访问以下五个地址验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://www.sse.com.cn/`
  - `https://finance.sina.com.cn/`（新浪财经）
  - `https://web.ifzq.gtimg.cn/`
  - `https://proxy.finance.qq.com/`
  每个代理分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次；**协议** 写入 meta：只通 HTTP→`http`，只通 HTTPS→`https`，两个都通→`http/s`。去重按 **协议+IP+端口**。校验时**多代理并发**、**单代理内五域名并行**。列表直显：表格列为「代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议」。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=4000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 4000，最大 4000。遇目标站限流可适当调低。

### 工作流程

1. **拉取**：从 `sources/` 目录下配置的源（仅处理 `http.txt`、`https.txt`）拉取原始代理数据，支持动态 URL 及 Base64 等格式。
2. **解析与规范化**：将原始数据解析为标准代理格式（协议、IP、端口、认证等）。
3. **校验**：对 HTTP/HTTPS 代理通过上述验证与 2 秒超时规则进行筛选。
4. **去重与存储**：通过校验的代理去重后写入内存。
5. **生成列表**：按协议生成 `list/` 目录下的 `http.txt`、`https.txt`，并更新统计与 README 中的下载表格。

自动化由 GitHub Actions 执行：**全量流程**（抓取→解析→验证→生成列表）**每 6 小时**运行一次；**轻量复测**（对已有列表做连通性复测、剔除失效代理）**每 1 小时**运行一次。全量任务最长运行 12 小时，超时才会取消。下表「最后更新」时间为 UTC 及 UTC+8。

### 支持的代理格式示例

| 类型 | 格式 | 示例 |
| :--- | :--- | :--- |
| **HTTP/HTTPS** | `http://ip:port` | `http://1.2.3.4:8080` |
| | `https://ip:port` | `https://1.2.3.4:8080` |
| | `http://user:pass@ip:port` | `http://user:pass@1.2.3.4:8080` |

---

## 🔗 直接下载链接

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。上述三个文件（http.txt、https.txt、passed.txt）同时发布至 [GitHub Releases (tag: lists)](https://github.com/feitianyul/free-proxy-list/releases/tag/lists)，每次运行覆盖同版本附件，可固定使用该 Release 的附件 URL。

<!-- BEGIN PROXY LIST -->

最后更新：2026-03-10 18:39:14 UTC（2026-03-11 02:39:14 UTC+8）

**代理总数：43**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 43 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 43 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 185.115.74.185:8080 | ✓ 1152ms | ✓ 1967ms | ✓ 1839ms | 否 | 否 | http |
| 47.77.193.180:1080 | 否 | 否 | ✓ 579ms | ✓ 1412ms | ✓ 789ms | http |
| 217.76.245.80:999 | ✓ 700ms | ✓ 1494ms | ✓ 1141ms | ✓ 1451ms | ✓ 1287ms | http |
| 45.136.131.47:8443 | ✓ 324ms | ✓ 808ms | ✓ 1275ms | ✓ 861ms | ✓ 681ms | http |
| 115.231.181.40:8128 | ✓ 1177ms | 否 | 否 | ✓ 1429ms | ✓ 1072ms | http |
| 185.109.21.39:3128 | ✓ 1078ms | 否 | ✓ 1112ms | ✓ 1698ms | ✓ 1502ms | http |
| 190.9.109.198:999 | 否 | ✓ 1453ms | ✓ 1156ms | ✓ 1454ms | ✓ 1141ms | http |
| 45.136.130.223:8443 | ✓ 1280ms | 否 | ✓ 1234ms | ✓ 1898ms | 否 | http |
| 35.225.22.61:80 | ✓ 477ms | ✓ 1847ms | ✓ 988ms | ✓ 1037ms | ✓ 742ms | http |
| 165.227.5.10:8888 | ✓ 900ms | ✓ 1541ms | 否 | 否 | ✓ 895ms | http |
| 91.107.141.42:8081 | ✓ 1646ms | 否 | ✓ 1294ms | ✓ 1716ms | ✓ 1568ms | http |
| 67.169.98.211:443 | ✓ 1253ms | 否 | ✓ 1658ms | 否 | ✓ 1297ms | http |
| 202.155.12.161:443 | ✓ 1837ms | 否 | 否 | ✓ 1950ms | ✓ 1171ms | http |
| 101.47.73.135:3128 | ✓ 1682ms | 否 | ✓ 1774ms | 否 | ✓ 1960ms | http |
| 114.55.226.123:10086 | ✓ 1132ms | ✓ 1537ms | ✓ 1136ms | ✓ 1455ms | ✓ 1211ms | http |
| 207.254.71.62:8088 | ✓ 1643ms | ✓ 1676ms | ✓ 800ms | 否 | 否 | http |
| 211.95.152.50:45046 | ✓ 1339ms | ✓ 1883ms | ✓ 1473ms | ✓ 1593ms | 否 | http |
| 128.199.247.154:3128 | ✓ 868ms | 否 | ✓ 1197ms | ✓ 1292ms | ✓ 1138ms | http |
| 101.43.255.96:80 | ✓ 1214ms | ✓ 1391ms | ✓ 1708ms | ✓ 1506ms | ✓ 1431ms | http |
| 107.172.125.217:3128 | ✓ 782ms | 否 | ✓ 897ms | ✓ 1430ms | ✓ 695ms | http |
| 59.46.216.131:30001 | ✓ 1391ms | ✓ 1493ms | ✓ 1223ms | 否 | ✓ 1215ms | http |
| 39.104.201.40:7890 | ✓ 1641ms | 否 | ✓ 1137ms | ✓ 1953ms | 否 | http |
| 111.79.111.126:3128 | 否 | ✓ 1865ms | ✓ 1535ms | ✓ 1686ms | 否 | http |
| 103.86.131.62:80 | ✓ 1617ms | 否 | 否 | ✓ 1900ms | ✓ 1495ms | http |
| 101.32.244.83:8080 | ✓ 1105ms | 否 | ✓ 1071ms | ✓ 1562ms | ✓ 1385ms | http |
| 121.43.196.213:8222 | ✓ 1062ms | ✓ 1236ms | ✓ 932ms | ✓ 1263ms | ✓ 1023ms | http |
| 121.43.196.210:8222 | ✓ 1071ms | ✓ 1176ms | ✓ 973ms | ✓ 1293ms | ✓ 1031ms | http |
| 81.70.169.194:80 | ✓ 1074ms | 否 | ✓ 1137ms | ✓ 1376ms | ✓ 1087ms | http |
| 45.136.198.40:3128 | ✓ 971ms | 否 | ✓ 951ms | ✓ 1844ms | ✓ 1646ms | http |
| 45.186.6.104:3128 | ✓ 1239ms | ✓ 1847ms | ✓ 1579ms | 否 | 否 | http |
| 138.124.90.140:1080 | ✓ 1076ms | ✓ 1890ms | ✓ 1446ms | ✓ 1835ms | ✓ 1230ms | http |
| 138.124.53.25:7443 | ✓ 449ms | 否 | ✓ 1813ms | ✓ 1549ms | ✓ 1425ms | http |
| 62.113.119.14:8080 | ✓ 1373ms | 否 | ✓ 1247ms | 否 | ✓ 1531ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1203ms | ✓ 1370ms | ✓ 1309ms | ✓ 1131ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1363ms | ✓ 1187ms | 否 | ✓ 1108ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1377ms | 否 | ✓ 1404ms | ✓ 1080ms | http |
| 47.105.98.23:3128 | ✓ 971ms | ✓ 1274ms | ✓ 1085ms | 否 | ✓ 1073ms | http |
| 34.101.184.164:3128 | ✓ 1796ms | 否 | ✓ 1733ms | 否 | ✓ 1290ms | http |
| 103.39.51.190:8080 | ✓ 1690ms | 否 | 否 | ✓ 1442ms | ✓ 1513ms | http |
| 46.183.25.8:443 | ✓ 1760ms | 否 | ✓ 1112ms | ✓ 1931ms | ✓ 1247ms | http |
| 8.219.97.248:80 | ✓ 1848ms | 否 | ✓ 1483ms | 否 | ✓ 1487ms | http |
| 212.175.29.184:8080 | ✓ 713ms | ✓ 1999ms | 否 | 否 | ✓ 1860ms | http |
| 61.52.131.172:8443 | ✓ 1024ms | ✓ 1241ms | ✓ 1048ms | 否 | ✓ 1036ms | http |

<!-- END PROXY TABLE -->

## 🤝 参与贡献

本项目由社区驱动，欢迎任何形式的贡献。最简单的参与方式就是添加新的代理数据源。

请先阅读 **[贡献指南](CONTRIBUTING.md)** 了解如何开始。

## 🙏 支持本项目

如果您觉得本项目有帮助，欢迎给予支持，让更多人看到并参与贡献。

-   在 GitHub 上 **给本仓库加星** ⭐️
-   **分享**给朋友和同事

## ⚠️ 免责声明

-   本仓库中的代理均来自公开来源，不保证其速度、安全性或可用性。
-   使用这些代理的风险由您自行承担。
-   本仓库维护者不对任何滥用行为负责。请勿将代理用于非法用途。

## 📝 许可证

本仓库采用 MIT 许可证发布。详见 [LICENSE](LICENSE)。

## Stars
[![Star History Chart](https://api.star-history.com/svg?repos=feitianyul/free-proxy-list&type=Date)](https://star-history.com/#feitianyul/free-proxy-list&Date)
