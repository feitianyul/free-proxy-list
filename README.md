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

最后更新：2026-04-11 17:37:39 UTC（2026-04-12 01:37:39 UTC+8）

**代理总数：47**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 47 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 47 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 873ms | 否 | ✓ 853ms | 否 | ✓ 1295ms | http |
| 167.103.115.102:8800 | ✓ 1039ms | 否 | ✓ 1142ms | ✓ 1228ms | ✓ 1204ms | http |
| 113.160.132.26:8080 | ✓ 1563ms | ✓ 1440ms | ✓ 1824ms | ✓ 1096ms | ✓ 883ms | http |
| 45.167.124.52:8080 | ✓ 1399ms | 否 | ✓ 1672ms | 否 | ✓ 1621ms | http |
| 104.129.202.127:12354 | ✓ 369ms | ✓ 672ms | ✓ 707ms | ✓ 703ms | ✓ 689ms | http |
| 104.129.202.127:10810 | ✓ 266ms | ✓ 710ms | ✓ 705ms | ✓ 701ms | ✓ 708ms | http |
| 121.130.199.80:3128 | ✓ 790ms | ✓ 1247ms | ✓ 1774ms | ✓ 1107ms | ✓ 870ms | http |
| 167.103.34.108:8800 | ✓ 1097ms | 否 | ✓ 1073ms | ✓ 1324ms | ✓ 1252ms | http |
| 167.103.144.127:8800 | ✓ 1259ms | 否 | ✓ 1163ms | 否 | ✓ 1412ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1369ms | ✓ 1610ms | ✓ 1842ms | http |
| 167.103.31.122:8800 | ✓ 1763ms | 否 | ✓ 1793ms | 否 | ✓ 1900ms | http |
| 45.167.125.21:999 | ✓ 1193ms | 否 | 否 | ✓ 1863ms | ✓ 1706ms | http |
| 218.108.131.186:17890 | 否 | ✓ 1292ms | ✓ 830ms | ✓ 1510ms | ✓ 1450ms | http |
| 101.43.127.100:8877 | ✓ 827ms | ✓ 985ms | 否 | ✓ 1558ms | ✓ 833ms | http |
| 147.161.239.240:8800 | ✓ 774ms | ✓ 1901ms | ✓ 1262ms | 否 | ✓ 1586ms | http |
| 35.225.22.61:80 | 否 | ✓ 1524ms | ✓ 1267ms | ✓ 1471ms | 否 | http |
| 46.30.46.133:3128 | ✓ 1730ms | 否 | ✓ 1508ms | ✓ 1998ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1063ms | ✓ 1364ms | ✓ 1143ms | 否 | ✓ 1048ms | http |
| 103.82.23.118:5234 | ✓ 1194ms | 否 | ✓ 1170ms | ✓ 1512ms | ✓ 1460ms | http |
| 103.242.105.76:80 | ✓ 1778ms | 否 | ✓ 1921ms | ✓ 1732ms | ✓ 1682ms | http |
| 121.230.9.113:1080 | ✓ 985ms | ✓ 1379ms | ✓ 1793ms | ✓ 1247ms | ✓ 985ms | http |
| 114.231.72.214:1080 | ✓ 1700ms | ✓ 1105ms | ✓ 914ms | ✓ 1949ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1210ms | ✓ 1421ms | ✓ 1567ms | ✓ 1369ms | ✓ 1165ms | http |
| 5.9.55.221:5000 | ✓ 1284ms | 否 | ✓ 1964ms | 否 | ✓ 1899ms | http |
| 162.240.154.26:3128 | ✓ 490ms | 否 | ✓ 1951ms | 否 | ✓ 723ms | http |
| 152.32.132.190:7890 | ✓ 1116ms | 否 | 否 | ✓ 858ms | ✓ 888ms | http |
| 120.92.212.16:7890 | ✓ 1045ms | ✓ 1609ms | ✓ 1219ms | 否 | 否 | http |
| 113.176.92.71:3128 | ✓ 1805ms | ✓ 1279ms | 否 | ✓ 1211ms | 否 | http |
| 121.230.9.203:1080 | ✓ 1359ms | ✓ 1327ms | ✓ 1215ms | ✓ 1284ms | ✓ 1241ms | http |
| 47.105.98.23:3128 | ✓ 1821ms | ✓ 1145ms | 否 | ✓ 1623ms | ✓ 955ms | http |
| 8.209.238.110:47701 | ✓ 461ms | ✓ 1789ms | ✓ 636ms | ✓ 828ms | ✓ 651ms | http |
| 115.231.181.40:8128 | ✓ 806ms | 否 | ✓ 1495ms | ✓ 1107ms | ✓ 1130ms | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 858ms | ✓ 1102ms | ✓ 900ms | http |
| 43.99.47.107:56615 | ✓ 718ms | ✓ 1108ms | ✓ 773ms | ✓ 762ms | ✓ 625ms | http |
| 207.254.71.62:8088 | ✓ 1256ms | 否 | ✓ 1905ms | 否 | ✓ 1912ms | http |
| 155.117.18.36:25388 | ✓ 1602ms | ✓ 1664ms | 否 | ✓ 1974ms | 否 | http |
| 110.42.37.202:20005 | 否 | ✓ 1389ms | ✓ 1606ms | ✓ 1435ms | ✓ 1197ms | http |
| 222.228.171.92:8080 | ✓ 1412ms | 否 | ✓ 602ms | ✓ 1645ms | 否 | http |
| 111.79.111.126:3128 | ✓ 1467ms | 否 | 否 | ✓ 1212ms | ✓ 1114ms | http |
| 159.223.225.118:8888 | ✓ 1635ms | 否 | 否 | ✓ 1604ms | ✓ 1816ms | http |
| 103.113.70.189:1082 | ✓ 544ms | ✓ 1783ms | 否 | ✓ 1451ms | 否 | http |
| 24.144.86.173:1080 | ✓ 292ms | ✓ 847ms | ✓ 986ms | ✓ 683ms | ✓ 510ms | http |
| 47.238.220.4:8888 | ✓ 1908ms | 否 | 否 | ✓ 1784ms | ✓ 990ms | http |
| 129.213.139.179:8080 | ✓ 961ms | ✓ 1638ms | ✓ 1261ms | 否 | ✓ 1519ms | http |
| 103.113.70.189:1081 | ✓ 375ms | ✓ 1394ms | ✓ 522ms | ✓ 1215ms | ✓ 1175ms | http |
| 202.141.161.53:10808 | ✓ 1335ms | ✓ 1485ms | 否 | 否 | ✓ 1141ms | http |
| 168.110.52.228:3128 | ✓ 1537ms | 否 | ✓ 1751ms | 否 | ✓ 1870ms | http |

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
