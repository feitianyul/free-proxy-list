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

最后更新：2026-03-15 03:41:22 UTC（2026-03-15 11:41:22 UTC+8）

**代理总数：54**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 53 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 54 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 113.160.132.26:8080 | ✓ 1989ms | 否 | ✓ 1215ms | 否 | ✓ 1221ms | http |
| 205.209.118.30:3138 | ✓ 1369ms | 否 | 否 | ✓ 1188ms | ✓ 1913ms | http |
| 120.92.212.16:7890 | ✓ 1018ms | 否 | ✓ 994ms | ✓ 1285ms | ✓ 1030ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1537ms | 否 | ✓ 1538ms | ✓ 1285ms | http |
| 45.140.147.82:1081 | ✓ 524ms | 否 | ✓ 1156ms | 否 | ✓ 1122ms | http |
| 35.225.22.61:80 | 否 | ✓ 1649ms | ✓ 198ms | 否 | ✓ 1146ms | http |
| 85.208.108.43:2094 | ✓ 343ms | 否 | ✓ 372ms | ✓ 1115ms | ✓ 684ms | http |
| 103.84.95.54:7890 | ✓ 1923ms | 否 | ✓ 798ms | ✓ 962ms | ✓ 718ms | http |
| 62.60.177.204:34094 | 否 | ✓ 1306ms | 否 | ✓ 1208ms | ✓ 795ms | http |
| 85.208.108.43:10808 | ✓ 347ms | 否 | ✓ 686ms | ✓ 1173ms | ✓ 789ms | http |
| 45.167.124.52:8080 | ✓ 1692ms | ✓ 1982ms | ✓ 1447ms | 否 | 否 | http |
| 210.223.44.230:3128 | ✓ 1803ms | ✓ 1226ms | ✓ 1186ms | ✓ 1149ms | 否 | http |
| 101.43.127.100:8877 | ✓ 946ms | ✓ 1264ms | ✓ 1570ms | 否 | ✓ 1692ms | http |
| 81.70.169.194:80 | ✓ 1150ms | ✓ 1358ms | ✓ 1086ms | 否 | 否 | http |
| 101.43.255.96:80 | 否 | ✓ 1674ms | ✓ 1289ms | 否 | ✓ 1102ms | http |
| 190.242.157.215:8080 | 否 | 否 | ✓ 1533ms | ✓ 1676ms | ✓ 1547ms | http |
| 45.88.0.99:3128 | ✓ 586ms | ✓ 1956ms | ✓ 1875ms | 否 | 否 | http |
| 20.210.39.153:8561 | 否 | ✓ 1103ms | ✓ 709ms | ✓ 935ms | ✓ 710ms | http |
| 38.145.203.135:8443 | 否 | 否 | ✓ 1803ms | ✓ 1654ms | ✓ 1157ms | http |
| 91.233.223.147:3128 | 否 | 否 | ✓ 1649ms | ✓ 1945ms | ✓ 1483ms | http |
| 20.78.26.206:8561 | ✓ 543ms | ✓ 1069ms | ✓ 544ms | ✓ 869ms | ✓ 736ms | http |
| 47.101.149.27:9010 | ✓ 1412ms | 否 | ✓ 1835ms | ✓ 1581ms | 否 | http |
| 85.198.96.242:3128 | ✓ 1067ms | 否 | ✓ 592ms | ✓ 1732ms | ✓ 1310ms | http |
| 138.124.53.221:443 | ✓ 501ms | 否 | ✓ 1865ms | ✓ 1861ms | ✓ 1211ms | http |
| 64.188.90.36:1080 | ✓ 620ms | ✓ 1703ms | ✓ 749ms | 否 | 否 | http |
| 45.88.0.117:3128 | 否 | ✓ 1855ms | 否 | ✓ 1453ms | ✓ 1272ms | http |
| 45.88.0.116:3128 | ✓ 1480ms | ✓ 1808ms | ✓ 1986ms | ✓ 1387ms | 否 | http |
| 45.88.0.111:3128 | 否 | ✓ 1758ms | ✓ 1519ms | ✓ 1398ms | ✓ 1773ms | http |
| 213.220.62.62:3128 | ✓ 1139ms | 否 | ✓ 1452ms | 否 | ✓ 1543ms | http |
| 165.227.5.10:8888 | ✓ 918ms | ✓ 1718ms | ✓ 1095ms | ✓ 918ms | 否 | http |
| 138.124.53.25:7443 | ✓ 989ms | 否 | ✓ 1334ms | ✓ 1858ms | ✓ 1253ms | http |
| 2.56.122.146:10808 | ✓ 945ms | 否 | ✓ 1206ms | ✓ 1619ms | ✓ 1120ms | http |
| 59.46.216.131:30001 | ✓ 1102ms | 否 | ✓ 1181ms | 否 | ✓ 1536ms | http |
| 45.136.130.163:8443 | ✓ 584ms | ✓ 1032ms | ✓ 334ms | ✓ 950ms | ✓ 712ms | http |
| 45.136.130.156:8443 | ✓ 582ms | ✓ 1050ms | ✓ 330ms | ✓ 961ms | ✓ 743ms | http |
| 45.136.130.161:8443 | ✓ 584ms | ✓ 1069ms | ✓ 319ms | ✓ 967ms | ✓ 726ms | http |
| 45.136.130.162:8443 | ✓ 582ms | ✓ 1076ms | ✓ 332ms | ✓ 989ms | ✓ 699ms | http |
| 45.136.130.155:8443 | ✓ 583ms | 否 | ✓ 308ms | ✓ 946ms | ✓ 758ms | http |
| 144.31.25.69:21064 | ✓ 1040ms | 否 | ✓ 1089ms | 否 | ✓ 1929ms | http |
| 88.80.150.82:8080 | ✓ 1747ms | ✓ 1976ms | 否 | 否 | ✓ 1804ms | https |
| 178.236.245.59:3128 | ✓ 752ms | 否 | ✓ 1388ms | ✓ 1749ms | ✓ 1339ms | http |
| 106.117.208.101:7890 | ✓ 1887ms | ✓ 1574ms | ✓ 1287ms | ✓ 1462ms | ✓ 1020ms | http |
| 38.145.203.161:8443 | ✓ 980ms | ✓ 1256ms | ✓ 326ms | ✓ 984ms | ✓ 738ms | http |
| 45.136.198.40:3128 | ✓ 736ms | ✓ 1604ms | ✓ 1233ms | ✓ 1477ms | ✓ 1459ms | http |
| 34.101.184.164:3128 | ✓ 1631ms | 否 | ✓ 1030ms | ✓ 1432ms | ✓ 1144ms | http |
| 95.3.9.78:3128 | 否 | 否 | ✓ 1837ms | ✓ 1756ms | ✓ 1338ms | http |
| 194.5.212.40:8080 | ✓ 1068ms | ✓ 1546ms | ✓ 1548ms | 否 | ✓ 1676ms | http |
| 146.19.128.135:1080 | ✓ 1075ms | 否 | ✓ 748ms | ✓ 1951ms | ✓ 1914ms | http |
| 61.52.131.172:8443 | ✓ 943ms | ✓ 1214ms | ✓ 1014ms | ✓ 1242ms | ✓ 954ms | http |
| 121.230.8.250:1080 | ✓ 1582ms | 否 | ✓ 1230ms | ✓ 1625ms | 否 | http |
| 103.113.70.189:1081 | ✓ 171ms | ✓ 1010ms | 否 | ✓ 1246ms | ✓ 942ms | http |
| 101.47.73.135:3128 | ✓ 993ms | 否 | ✓ 1077ms | ✓ 1155ms | ✓ 1185ms | http |
| 211.171.114.154:3128 | 否 | 否 | ✓ 1428ms | ✓ 1554ms | ✓ 1702ms | http |
| 95.3.9.78:8080 | ✓ 824ms | 否 | ✓ 813ms | 否 | ✓ 1485ms | http |

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
