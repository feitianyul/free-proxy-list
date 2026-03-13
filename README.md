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

最后更新：2026-03-13 15:36:16 UTC（2026-03-13 23:36:16 UTC+8）

**代理总数：37**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 36 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 37 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.136.130.175:8443 | ✓ 642ms | 否 | ✓ 972ms | ✓ 928ms | ✓ 620ms | http |
| 205.209.118.30:3138 | ✓ 210ms | ✓ 1813ms | ✓ 1258ms | ✓ 1223ms | ✓ 1011ms | http |
| 62.60.177.204:34094 | ✓ 327ms | 否 | ✓ 968ms | ✓ 994ms | ✓ 1591ms | http |
| 103.84.95.54:7890 | ✓ 944ms | 否 | ✓ 1159ms | ✓ 1270ms | ✓ 1408ms | http |
| 152.42.213.210:8080 | ✓ 1679ms | 否 | ✓ 1810ms | ✓ 1180ms | ✓ 1237ms | http |
| 152.42.213.210:443 | ✓ 1679ms | 否 | ✓ 1811ms | ✓ 1755ms | ✓ 1289ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1964ms | ✓ 1367ms | ✓ 1675ms | ✓ 1092ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 967ms | ✓ 1229ms | ✓ 901ms | http |
| 103.113.70.189:1081 | 否 | 否 | ✓ 1134ms | ✓ 1200ms | ✓ 902ms | http |
| 178.236.245.17:3128 | ✓ 913ms | ✓ 1966ms | ✓ 1502ms | 否 | 否 | http |
| 120.92.212.16:8890 | ✓ 1079ms | ✓ 1970ms | ✓ 1005ms | ✓ 1856ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1082ms | 否 | ✓ 1282ms | 否 | ✓ 1975ms | http |
| 45.167.124.52:8080 | 否 | 否 | ✓ 1598ms | ✓ 1750ms | ✓ 1284ms | http |
| 81.70.169.194:80 | ✓ 1005ms | ✓ 1624ms | ✓ 1039ms | ✓ 1301ms | 否 | http |
| 14.225.212.37:7890 | ✓ 1496ms | ✓ 1490ms | 否 | ✓ 1921ms | ✓ 1011ms | http |
| 101.43.255.96:80 | ✓ 981ms | ✓ 1428ms | 否 | 否 | ✓ 1080ms | http |
| 45.136.130.223:8443 | 否 | 否 | ✓ 265ms | ✓ 816ms | ✓ 1105ms | http |
| 143.189.3.198:8080 | 否 | ✓ 992ms | ✓ 616ms | ✓ 958ms | ✓ 1723ms | http |
| 116.80.49.165:3172 | ✓ 1603ms | 否 | 否 | ✓ 1914ms | ✓ 1706ms | http |
| 120.92.212.16:7890 | ✓ 1040ms | 否 | ✓ 1038ms | ✓ 1668ms | 否 | http |
| 49.51.244.112:8888 | ✓ 688ms | ✓ 1594ms | ✓ 655ms | ✓ 1684ms | 否 | http |
| 47.77.193.180:1080 | 否 | ✓ 1519ms | ✓ 182ms | ✓ 797ms | ✓ 600ms | http |
| 86.53.183.16:1080 | ✓ 1257ms | 否 | ✓ 1231ms | 否 | ✓ 1476ms | http |
| 88.80.150.82:8080 | ✓ 1282ms | 否 | ✓ 1245ms | 否 | ✓ 1925ms | https |
| 119.92.142.80:8082 | ✓ 1612ms | 否 | ✓ 1901ms | ✓ 1498ms | 否 | http |
| 192.71.213.85:9090 | ✓ 1034ms | 否 | ✓ 979ms | ✓ 1385ms | 否 | http |
| 45.136.198.40:3128 | ✓ 796ms | ✓ 1578ms | ✓ 1037ms | ✓ 1580ms | ✓ 1291ms | http |
| 101.47.73.135:3128 | ✓ 1536ms | 否 | ✓ 1269ms | 否 | ✓ 1226ms | http |
| 198.24.188.138:37835 | ✓ 1131ms | 否 | 否 | ✓ 1850ms | ✓ 1088ms | http |
| 103.39.51.190:8080 | ✓ 1962ms | 否 | 否 | ✓ 1706ms | ✓ 1507ms | http |
| 14.225.211.139:7890 | 否 | ✓ 1676ms | 否 | ✓ 1778ms | ✓ 1250ms | http |
| 103.87.67.75:3129 | ✓ 1824ms | 否 | ✓ 1940ms | 否 | ✓ 1243ms | http |
| 61.52.131.172:8443 | ✓ 925ms | ✓ 1219ms | ✓ 1007ms | ✓ 1242ms | ✓ 1022ms | http |
| 172.212.68.37:3128 | ✓ 1171ms | 否 | ✓ 1121ms | ✓ 1265ms | ✓ 1017ms | http |
| 180.127.149.247:1080 | ✓ 1659ms | 否 | 否 | ✓ 1403ms | ✓ 1667ms | http |
| 159.223.42.219:3128 | ✓ 1506ms | 否 | ✓ 1877ms | ✓ 1545ms | ✓ 1483ms | http |
| 45.140.147.82:1081 | ✓ 589ms | 否 | ✓ 1491ms | ✓ 1903ms | 否 | http |

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
