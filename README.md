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

最后更新：2026-05-01 16:55:23 UTC（2026-05-02 00:55:23 UTC+8）

**代理总数：60**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 60 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 60 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 47.85.51.197:1080 | ✓ 190ms | ✓ 1101ms | ✓ 403ms | 否 | 否 | http |
| 34.96.238.40:8080 | 否 | ✓ 1481ms | ✓ 1148ms | ✓ 1351ms | 否 | http |
| 206.206.126.177:2412 | ✓ 731ms | 否 | ✓ 1301ms | ✓ 1017ms | ✓ 798ms | http |
| 1.231.81.166:3128 | ✓ 1053ms | ✓ 1453ms | 否 | ✓ 1045ms | ✓ 1007ms | http |
| 46.101.95.183:8888 | ✓ 1797ms | 否 | ✓ 1224ms | ✓ 1830ms | ✓ 1456ms | http |
| 113.160.132.26:8080 | 否 | 否 | ✓ 952ms | ✓ 1595ms | ✓ 1111ms | http |
| 223.84.151.86:30005 | ✓ 1993ms | 否 | ✓ 1916ms | ✓ 1685ms | ✓ 1839ms | http |
| 45.167.124.71:999 | ✓ 1099ms | ✓ 1874ms | ✓ 1381ms | 否 | ✓ 1987ms | http |
| 103.35.190.69:1081 | ✓ 512ms | ✓ 1224ms | ✓ 833ms | ✓ 1567ms | ✓ 939ms | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1395ms | ✓ 1159ms | ✓ 1826ms | http |
| 217.76.245.80:999 | ✓ 1104ms | ✓ 1674ms | ✓ 1146ms | ✓ 1501ms | ✓ 1227ms | http |
| 218.108.131.186:17890 | ✓ 829ms | ✓ 1073ms | ✓ 884ms | ✓ 1094ms | ✓ 906ms | http |
| 120.92.212.16:7890 | ✓ 999ms | ✓ 1418ms | ✓ 1073ms | 否 | 否 | http |
| 20.78.118.91:8561 | ✓ 1358ms | ✓ 1024ms | ✓ 668ms | ✓ 861ms | 否 | http |
| 20.210.39.153:8561 | ✓ 1370ms | ✓ 1172ms | ✓ 583ms | ✓ 858ms | 否 | http |
| 20.78.26.206:8561 | ✓ 1366ms | 否 | ✓ 580ms | 否 | ✓ 722ms | http |
| 86.104.72.220:1081 | ✓ 496ms | 否 | ✓ 429ms | ✓ 1321ms | ✓ 995ms | http |
| 148.230.4.241:999 | ✓ 1208ms | ✓ 1342ms | ✓ 575ms | ✓ 1579ms | ✓ 1394ms | http |
| 43.133.44.89:8888 | ✓ 1529ms | 否 | ✓ 738ms | 否 | ✓ 1209ms | http |
| 158.178.237.243:3128 | ✓ 1488ms | 否 | ✓ 1966ms | ✓ 1428ms | ✓ 1106ms | http |
| 103.157.200.126:3128 | ✓ 1612ms | 否 | 否 | ✓ 1925ms | ✓ 1515ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1237ms | ✓ 1032ms | ✓ 1519ms | 否 | http |
| 8.219.97.248:80 | ✓ 1196ms | 否 | ✓ 1543ms | 否 | ✓ 1685ms | http |
| 62.113.119.14:8080 | ✓ 823ms | 否 | ✓ 1329ms | ✓ 1738ms | 否 | http |
| 20.27.14.220:8561 | 否 | 否 | ✓ 714ms | ✓ 825ms | ✓ 623ms | http |
| 20.27.15.111:8561 | 否 | 否 | ✓ 713ms | ✓ 824ms | ✓ 629ms | http |
| 20.27.11.248:8561 | 否 | 否 | ✓ 719ms | ✓ 822ms | ✓ 629ms | http |
| 20.27.13.35:8561 | 否 | 否 | ✓ 716ms | ✓ 841ms | ✓ 624ms | http |
| 86.104.72.220:1082 | ✓ 1108ms | ✓ 1924ms | ✓ 959ms | 否 | ✓ 876ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1396ms | ✓ 1078ms | 否 | ✓ 1424ms | http |
| 94.131.118.129:1082 | ✓ 1179ms | ✓ 1261ms | ✓ 1004ms | 否 | ✓ 1289ms | http |
| 91.184.241.12:443 | ✓ 841ms | 否 | ✓ 1088ms | 否 | ✓ 1824ms | http |
| 89.208.106.138:10808 | ✓ 1679ms | 否 | ✓ 649ms | 否 | ✓ 1494ms | http |
| 101.32.244.83:8080 | ✓ 998ms | 否 | ✓ 960ms | ✓ 1271ms | ✓ 1238ms | http |
| 121.43.196.213:8222 | ✓ 924ms | ✓ 1088ms | ✓ 864ms | ✓ 1173ms | ✓ 963ms | http |
| 121.43.196.210:8222 | ✓ 1039ms | ✓ 1058ms | ✓ 891ms | ✓ 1167ms | ✓ 956ms | http |
| 152.32.132.190:7890 | ✓ 856ms | 否 | ✓ 912ms | 否 | ✓ 844ms | http |
| 8.154.21.175:3128 | ✓ 870ms | ✓ 1088ms | ✓ 940ms | ✓ 1124ms | ✓ 951ms | http |
| 101.32.243.189:80 | ✓ 957ms | ✓ 1676ms | 否 | ✓ 1423ms | ✓ 1267ms | http |
| 80.92.204.47:1081 | ✓ 1256ms | ✓ 1553ms | ✓ 1097ms | ✓ 1783ms | ✓ 1154ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1241ms | ✓ 1673ms | ✓ 1420ms | http |
| 103.35.190.69:1082 | ✓ 1491ms | 否 | ✓ 681ms | ✓ 1646ms | 否 | http |
| 183.238.3.150:7897 | ✓ 921ms | ✓ 1255ms | ✓ 1050ms | ✓ 1100ms | ✓ 891ms | http |
| 183.232.248.73:7890 | ✓ 878ms | ✓ 1198ms | ✓ 900ms | ✓ 1075ms | ✓ 911ms | http |
| 220.197.44.36:3128 | ✓ 1526ms | ✓ 1505ms | ✓ 1501ms | 否 | ✓ 1411ms | http |
| 3.101.133.120:80 | ✓ 1118ms | 否 | ✓ 1550ms | ✓ 1305ms | ✓ 970ms | http |
| 120.92.108.86:7890 | ✓ 1372ms | 否 | 否 | ✓ 1855ms | ✓ 1549ms | http |
| 150.107.140.238:3128 | ✓ 1509ms | 否 | ✓ 948ms | 否 | ✓ 964ms | http |
| 51.95.13.205:2514 | ✓ 899ms | 否 | ✓ 1230ms | 否 | ✓ 1799ms | http |
| 72.11.151.159:6005 | ✓ 1232ms | ✓ 1849ms | ✓ 934ms | ✓ 1686ms | ✓ 1230ms | http |
| 91.217.81.131:1080 | ✓ 1813ms | 否 | ✓ 906ms | 否 | ✓ 1902ms | http |
| 129.213.139.179:8080 | ✓ 438ms | 否 | ✓ 680ms | ✓ 1363ms | ✓ 1051ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1039ms | ✓ 1817ms | 否 | ✓ 735ms | http |
| 130.61.174.200:1080 | ✓ 1075ms | ✓ 1876ms | ✓ 752ms | 否 | ✓ 1632ms | http |
| 182.53.202.208:8080 | ✓ 1334ms | 否 | ✓ 1430ms | ✓ 1723ms | ✓ 1338ms | http |
| 116.63.160.98:8899 | ✓ 1045ms | ✓ 1229ms | ✓ 1051ms | ✓ 1358ms | ✓ 1057ms | http |
| 61.52.131.172:8443 | ✓ 864ms | ✓ 1190ms | ✓ 1023ms | ✓ 1252ms | ✓ 983ms | http |
| 16.18.37.186:53208 | ✓ 1043ms | 否 | ✓ 1223ms | 否 | ✓ 1906ms | http |
| 57.128.188.167:9157 | ✓ 1702ms | 否 | ✓ 1634ms | 否 | ✓ 1826ms | http |
| 94.131.118.129:1081 | 否 | ✓ 1814ms | ✓ 1198ms | 否 | ✓ 1423ms | http |

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
