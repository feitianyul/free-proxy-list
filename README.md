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

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。

<!-- BEGIN PROXY LIST -->

最后更新：2026-02-22 10:44:00 UTC（2026-02-22 18:44:00 UTC+8）

**代理总数：20**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 20 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 20 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.72.254.71:3128 | ✓ 1605ms | ✓ 993ms | ✓ 525ms | ✓ 1311ms | ✓ 903ms | http |
| 211.230.49.122:3128 | ✓ 1628ms | ✓ 867ms | ✓ 1859ms | ✓ 1364ms | ✓ 1091ms | http |
| 186.148.180.46:999 | ✓ 1243ms | ✓ 1932ms | ✓ 1613ms | ✓ 1823ms | ✓ 1518ms | http |
| 120.92.212.16:7890 | ✓ 1607ms | ✓ 1287ms | ✓ 1466ms | ✓ 1484ms | ✓ 1336ms | http |
| 51.81.46.174:3128 | ✓ 1056ms | ✓ 1939ms | ✓ 1489ms | ✓ 1510ms | ✓ 1096ms | http |
| 120.92.212.16:8890 | ✓ 973ms | ✓ 1404ms | ✓ 1117ms | ✓ 1259ms | ✓ 1740ms | http |
| 52.54.20.49:80 | ✓ 390ms | ✓ 1706ms | ✓ 1592ms | ✓ 1467ms | ✓ 1862ms | http |
| 195.123.209.48:3128 | ✓ 806ms | ✓ 1653ms | ✓ 876ms | ✓ 1692ms | ✓ 1980ms | http |
| 54.173.56.172:80 | ✓ 295ms | ✓ 1261ms | ✓ 1240ms | ✓ 1587ms | ✓ 1551ms | http |
| 47.122.124.35:7890 | ✓ 1247ms | ✓ 1300ms | ✓ 1148ms | ✓ 1186ms | ✓ 976ms | http |
| 183.98.143.134:8077 | ✓ 1009ms | ✓ 1842ms | ✓ 840ms | ✓ 1559ms | ✓ 774ms | http |
| 205.209.118.30:3138 | ✓ 334ms | ✓ 1181ms | ✓ 1981ms | ✓ 1895ms | ✓ 1031ms | http |
| 103.215.36.88:17224 | ✓ 944ms | ✓ 1148ms | ✓ 1335ms | ✓ 1292ms | ✓ 1315ms | http |
| 36.147.78.166:80 | ✓ 1572ms | ✓ 1573ms | ✓ 1599ms | ✓ 1786ms | ✓ 1419ms | http |
| 35.234.17.221:8080 | ✓ 1325ms | ✓ 1049ms | ✓ 1274ms | ✓ 940ms | ✓ 965ms | http |
| 190.9.109.198:999 | ✓ 989ms | ✓ 1543ms | ✓ 1363ms | ✓ 1513ms | ✓ 1423ms | http |
| 137.220.150.22:6005 | ✓ 863ms | ✓ 1944ms | ✓ 929ms | ✓ 1141ms | ✓ 883ms | http |
| 36.136.27.2:4999 | ✓ 1167ms | ✓ 1259ms | ✓ 1188ms | ✓ 1341ms | ✓ 1102ms | http |
| 89.208.85.78:443 | ✓ 1141ms | ✓ 1875ms | ✓ 786ms | ✓ 1957ms | ✓ 1851ms | http |
| 45.140.147.155:1082 | ✓ 606ms | ✓ 1996ms | ✓ 658ms | ✓ 1888ms | ✓ 1280ms | http |

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
