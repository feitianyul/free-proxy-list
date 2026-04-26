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

最后更新：2026-04-26 15:40:31 UTC（2026-04-26 23:40:31 UTC+8）

**代理总数：30**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 30 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 30 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 20.120.225.109:3128 | ✓ 509ms | ✓ 1304ms | ✓ 1088ms | ✓ 1255ms | ✓ 1120ms | http |
| 217.76.245.80:999 | 否 | 否 | ✓ 1356ms | ✓ 1333ms | ✓ 1715ms | http |
| 211.95.152.50:45046 | ✓ 1187ms | ✓ 1543ms | 否 | ✓ 1984ms | 否 | http |
| 36.141.21.200:7890 | ✓ 1024ms | ✓ 1286ms | ✓ 1071ms | ✓ 1399ms | ✓ 1068ms | http |
| 47.85.51.197:1080 | ✓ 1817ms | 否 | ✓ 349ms | ✓ 1786ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1778ms | ✓ 1816ms | 否 | ✓ 1336ms | 否 | http |
| 20.78.118.91:8561 | ✓ 694ms | ✓ 1375ms | ✓ 979ms | ✓ 1249ms | ✓ 1339ms | http |
| 20.210.39.153:8561 | ✓ 709ms | 否 | ✓ 904ms | ✓ 1260ms | ✓ 1215ms | http |
| 20.78.26.206:8561 | ✓ 703ms | 否 | ✓ 916ms | ✓ 1346ms | ✓ 1265ms | http |
| 80.92.204.47:1081 | ✓ 1307ms | ✓ 1340ms | ✓ 491ms | 否 | ✓ 1534ms | http |
| 113.160.132.26:8080 | ✓ 1592ms | ✓ 1535ms | ✓ 1091ms | ✓ 1405ms | 否 | http |
| 43.133.90.161:8888 | ✓ 1671ms | 否 | ✓ 849ms | ✓ 1186ms | 否 | http |
| 46.101.95.183:8888 | ✓ 1397ms | 否 | ✓ 580ms | ✓ 1820ms | ✓ 1802ms | http |
| 101.32.244.83:8080 | ✓ 1920ms | 否 | ✓ 1090ms | ✓ 1556ms | ✓ 1415ms | http |
| 121.43.196.213:8222 | ✓ 1085ms | ✓ 1274ms | ✓ 993ms | ✓ 1258ms | 否 | http |
| 121.43.196.210:8222 | ✓ 1067ms | ✓ 1335ms | ✓ 1035ms | ✓ 1228ms | ✓ 1074ms | http |
| 1.231.81.166:3128 | 否 | 否 | ✓ 1677ms | ✓ 1748ms | ✓ 1841ms | http |
| 103.157.200.126:3128 | ✓ 1403ms | 否 | ✓ 1339ms | ✓ 1672ms | 否 | http |
| 34.101.184.164:3128 | ✓ 1822ms | 否 | ✓ 1058ms | ✓ 1481ms | ✓ 1544ms | http |
| 120.92.108.86:7890 | ✓ 1420ms | 否 | ✓ 1774ms | ✓ 1787ms | ✓ 1444ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1865ms | ✓ 1974ms | ✓ 1610ms | ✓ 1464ms | http |
| 45.59.122.132:80 | ✓ 1336ms | 否 | ✓ 1514ms | 否 | ✓ 1475ms | http |
| 62.113.119.14:8080 | ✓ 1019ms | 否 | ✓ 1047ms | ✓ 1537ms | ✓ 1135ms | http |
| 218.108.131.186:17890 | 否 | 否 | ✓ 1865ms | ✓ 1656ms | ✓ 1018ms | http |
| 38.180.192.119:3128 | ✓ 987ms | ✓ 1945ms | ✓ 1495ms | ✓ 1219ms | ✓ 1168ms | http |
| 47.105.98.23:3128 | ✓ 1020ms | ✓ 1889ms | 否 | ✓ 1673ms | ✓ 1096ms | http |
| 47.110.42.192:9003 | ✓ 1871ms | ✓ 1409ms | 否 | ✓ 1769ms | ✓ 1522ms | http |
| 86.104.72.220:1081 | ✓ 693ms | ✓ 1508ms | ✓ 130ms | ✓ 1029ms | ✓ 763ms | http |
| 61.52.131.172:8443 | ✓ 1060ms | ✓ 1336ms | ✓ 1031ms | 否 | 否 | http |
| 103.39.51.207:8080 | ✓ 1434ms | 否 | 否 | ✓ 1607ms | ✓ 1780ms | http |

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
